// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func main() {
	// Create a new context.
	ctx := cuecontext.New()

	// https://cuelang.org/docs/tour/types/closed/
	// A closed struct is a struct that has a fixed set of fields.
	schema := ctx.CompileString(`close({
		age: >=13
		name: string 
		hobby: [...string]
	})`)
	// hobby: [string] means hobby is a list of string, but only 1 item
	// hobby: [string, string] means hobby is a list of string, but only 2 items
	// hobby: [...string] means hobby is a list of indefinite string

	value := ctx.CompileString(`{
		"age": 13, 
		"name": "john",
		"hobby": ["reading", "coding"]
	}`)
	// Compile the schema and value.
	// Unify the value with the schema.
	// Validate the unified value.
	// Concrete(true) means the value must be concrete (required).
	fmt.Println("JSON", schema.Unify(value).Validate(
		cue.Concrete(true),
	))

	// Works with yaml?
	value = ctx.CompileString(`
	age: 13
	name: "john"
	hobby: ["reading", "coding"]
	`)
	fmt.Println("YAML", schema.Unify(value).Validate(
		cue.Concrete(true),
	))
}
