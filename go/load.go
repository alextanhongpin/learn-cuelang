// This example demonstrates the usage of multiple definitions.
package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func main() {
	instances := load.Instances([]string{
		"author.cue",
		"book.cue",
	}, nil)

	ctx := cuecontext.New()
	values, err := ctx.BuildInstances(instances)
	if err != nil {
		panic(err)
	}

	var unified cue.Value
	for _, v := range values {
		unified = unified.Unify(v)
	}

	schema := ctx.CompileString(`#Book`, cue.Scope(unified))
	if err := schema.Err(); err != nil {
		panic(err)
	}

	data := `{
	"id": "123",
	"year": 2024,
	"title": "My Book",
	"author": {
		"id": "456",
		"name": "John Appleseed"
	}
}`

	value := ctx.CompileString(data)
	if err := value.Err(); err != nil {
		panic(err)
	}
	value = value.Unify(schema)
	fmt.Println(value)
	fmt.Println(value.Validate(
		cue.Concrete(true),
	))
}
