// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cuego"
)

const rawSchema = `close({
		age!: >=13
		name!: string 
		#bytes: len(name) & >=1
		hobby!: [...string]
	})`

func main() {
	// Create a new context.
	ctx := cuecontext.New()

	// https://cuelang.org/docs/tour/types/closed/
	// A closed struct is a struct that has a fixed set of fields.
	schema := ctx.CompileString(rawSchema)
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

	// Unfortunately strings needs to be quoted.
	value = ctx.CompileString(`
	age: 13
	name: john
	hobby: ["reading", "coding"]
	`)
	fmt.Println("YAML", schema.Unify(value).Validate(
		cue.Concrete(true),
	))

	fmt.Println()
	fmt.Println("=== Testing struct")
	type User struct {
		Age   int      `json:"age" cue:">=13"`
		Name  string   `json:"name"`
		Hobby []string `json:"hobby"`
	}
	u := User{
		Age:   12,
		Hobby: []string{"reading", "coding"},
	}
	if err := cuego.Validate(u); err != nil {
		fmt.Println(fmt.Errorf("invalid user: %w", err))
	}
	fmt.Println(u)

	// For more complicated schema, use the raw schema.
	fmt.Println()
	fmt.Println("=== Testing constraint")
	// Register the constraint
	if err := cuego.Constrain(&User{}, rawSchema); err != nil {
		panic(err)
	}
	u.Age = 13
	u.Name = ""
	if err := cuego.Validate(&u); err != nil {
		fmt.Println(fmt.Errorf("invalid user: %w", err))
		fmt.Println(errors.Errors(err))
	}
	fmt.Println(u)
}
