package main

import (
	"strings"

	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cuego"
)

// User represents a user object.
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age" cue:">0 & <100"`
}

// Validate the user.
func (u *User) Validate() error {
	return cuego.Validate(u)
}

func main() {
	u := &User{Age: -1}
	if err := u.Validate(); err != nil {
		panic(errMsg(err))
	}
}

func errMsg(err error) string {
	a := []string{}
	for _, err := range errors.Errors(err) {
		a = append(a, err.Error())
	}
	s := strings.Join(a, "\n")
	if s == "" {
		return "nil"
	}
	return s
}
