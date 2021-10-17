package main

import (
	"fmt"
	"strings"

	_ "embed"

	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cuego"
)

//go:embed user.cue
var usercue string

func main() {
	type Config struct {
		Filename string
		OptFile  string `json:",omitempty"`
		MaxCount int
		MinCount int
	}

	err := cuego.Constrain(&Config{}, `{
		let jsonFile = =~".json$"

		// Filename must be defined and have a .json extension.
		Filename: jsonFile

		// Optfile must be undefined or be a file name with a .json extension
		OptFile?: jsonFile

		MinCount: >0 & <= MaxCount
		MaxCount: <= 10_000
	}`)
	fmt.Println("error:", errMsg(err))

	fmt.Println("validate:", errMsg(cuego.Validate(&Config{
		Filename: "foo.json",
		MaxCount: 1200,
		MinCount: 39,
	})))
	fmt.Println("")

	fmt.Println("validate:", errMsg(cuego.Validate(&Config{
		Filename: "foo.json",
		MaxCount: 12,
		MinCount: 39,
	})))
	fmt.Println("")

	fmt.Println("validate:", errMsg(cuego.Validate(&Config{
		Filename: "foo.jso",
		MaxCount: 120,
		MinCount: 39,
	})))
	fmt.Println("")

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err = cuego.Constrain(&User{}, usercue)
	fmt.Println("error:", errMsg(err))

	fmt.Println("validate:", errMsg(cuego.Validate(&User{
		Name: "John Doe",
		Age:  50,
	})))
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
