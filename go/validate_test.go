package main

import (
	"errors"
	"testing"

	"cuelang.org/go/cuego"
)

type User struct {
	Age   int      `json:"age"`
	Name  string   `json:"name"`
	Hobby []string `json:"hobby"`
}

const schema = `close({
		age!: >=13
		name!: string 
		#bytes: len(name) & >=1
		hobby!: [...string]
	})`

func init() {
	cuego.Constrain(&User{}, schema)
}

var user = User{
	Age:   12,
	Name:  "john",
	Hobby: []string{"reading", "coding"},
}

func (u *User) Validate() error {
	if u.Age < 13 {
		return errors.New("age must be greater than or equal to 13")
	}
	if len(u.Name) < 1 {
		return errors.New("name must have at least 1 character")
	}
	return nil
}

func BenchmarkCueValidate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = cuego.Validate(&user)
	}
}

func BenchmarkGoValidate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = user.Validate()
	}
}

// OUTPUT:
//➜  $ gotest -v -bench=. validate_test.go
//goos: darwin
//goarch: amd64
//cpu: Intel(R) Core(TM) i5-6267U CPU @ 2.90GHz
//BenchmarkCueValidate
//BenchmarkCueValidate-4             13542             84144 ns/op
//BenchmarkGoValidate
//BenchmarkGoValidate-4           1000000000               0.5671 ns/op
