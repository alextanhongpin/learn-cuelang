// This example demonstrates how to evaluate cue in golang.
package main

import (
	"fmt"

	"cuelang.org/go/cue"
)

const config = `
msg: "Hello, \(place)!"
place: string | *"world" // world is the default
`

func main() {
	var r cue.Runtime
	instance, err := r.Compile("config", config)
	if err != nil {
		panic(err)
	}

	str, err := instance.Lookup("msg").String()
	if err != nil {
		panic(err)
	}
	fmt.Println(str) // Hello, world!
	// We can fill the instance value.
	newInstance, err := instance.Fill("Malaysia", "place")
	if err != nil {
		panic(err)
	}
	fmt.Println(newInstance.Lookup("msg").String())

	// We can unify both values to get a new value.
	instance2, err := r.Compile("config", `{"place": "Malaysia"}`)
	if err != nil {
		panic(err)
	}

	val := instance2.Value().Unify(instance.Value())
	str, err = val.Lookup("msg").String()
	if err != nil {
		panic(err)
	}

	fmt.Println(str) // Hello, Malaysia!
}
