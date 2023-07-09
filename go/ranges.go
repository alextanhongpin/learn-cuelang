package main

import (
	"errors"
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

const module = `package ranges

min?: *0 | number // 0 if undefined
max?: number & >min // must be strictly greater than min if defined.
`

const sample = `{
	"min": 10,
	"max": -10
}`

func main() {
	// Same as:
	// $ cue vet ranges.json ranges.cue

	fmt.Println(method1())
	fmt.Println(method2())
}

func method1() error {
	ctx := cuecontext.New()
	schema := ctx.CompileString(module, cue.Filename("ranges.cue"))
	if err := schema.Err(); err != nil {
		return err
	}

	data := ctx.CompileString(sample, cue.Filename("range.json"))
	prog := schema.Unify(data)
	opts := []cue.Option{
		cue.Attributes(true),
		cue.Definitions(true),
		cue.Hidden(true),
	}

	err := prog.Validate(opts...)
	return err
}

func method2() error {
	ctx := cuecontext.New()
	binst := load.Instances([]string{
		"ranges.cue",
	}, nil)

	insts, err := ctx.BuildInstances(binst)
	if err != nil {
		return err
	}
	if len(insts) != 1 {
		return errors.New("more than 1 instance")
	}

	v := insts[0]
	data := ctx.CompileString(sample, cue.Filename("range.json"))
	unified := data.Unify(v)

	opts := []cue.Option{
		cue.Attributes(true),
		cue.Definitions(true),
		cue.Hidden(true),
	}

	err = unified.Validate(opts...)
	return err
}
