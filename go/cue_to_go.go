package main

import (
	"fmt"

	"cuelang.org/go/cue"
)

type ab struct {
	A, B int
}

type cd struct {
	C, D string
}

const data = `
A: 1
B: 2
B: >A
`

func main() {
	var r cue.Runtime
	instance, err := r.Compile("data", data)
	if err != nil {
		panic(err)
	}

	var one ab
	err = instance.Value().Decode(&one)
	if err != nil {
		panic(err)
	}
	fmt.Println(one)

	newInstance, err := instance.Fill(-1, "B")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", newInstance.Value())
	var two ab
	if err := newInstance.Value().Decode(&two); err != nil {
		fmt.Println(err)
	}
	fmt.Println(two)

	// Somehow this is false
	var three cd
	err = instance.Value().Decode(&three)
	if err != nil {
		panic(err)
	}
	fmt.Println(three)

	//b, err := gocode.Generate("github.com/alextanhongpin/learn-cuelang/go", instance, nil)
	//if err != nil {
	//panic(err)
	//}

	//err = os.WriteFile("cue_gen.go", b, 0644)
}
