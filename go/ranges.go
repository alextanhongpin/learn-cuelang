package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
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
	fmt.Println(method3())
	fmt.Println(method4())
	fmt.Println(method5())
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

	v := ctx.BuildInstance(binst[0])
	if err := v.Err(); err != nil {
		panic(err)
	}

	data := ctx.CompileString(sample, cue.Filename("range.json"))
	unified := data.Unify(v)

	opts := []cue.Option{
		cue.Attributes(true),
		cue.Definitions(true),
		cue.Hidden(true),
	}

	err := unified.Validate(opts...)
	return err
}

func method3() error {
	const module = `package ranges

#Range: {
	min?: *0 | number // 0 if undefined
	max?: number & >min // must be strictly greater than min if defined.
}
`
	const sample = `#Range & {
	"min": 10,
	"max": -10
}`
	ctx := cuecontext.New()
	schema := ctx.CompileString(module, cue.Filename("ranges.cue"))
	if err := schema.Err(); err != nil {
		return err
	}

	prog := ctx.CompileString(sample, cue.Scope(schema), cue.Filename("range.json"))
	opts := []cue.Option{
		cue.Attributes(true),
		cue.Definitions(true),
		cue.Hidden(true),
	}

	err := prog.Validate(opts...)
	return err
}

func method4() error {
	const sample = `#Range & {
	"min": 10,
	"max": -10
}`
	ctx := cuecontext.New()
	bins := load.Instances([]string{"ranges.v2.cue"}, nil)
	schema := ctx.BuildInstance(bins[0])

	prog := ctx.CompileString(sample, cue.Scope(schema), cue.Filename("range.json"))
	opts := []cue.Option{
		cue.Attributes(true),
		cue.Definitions(true),
		cue.Hidden(true),
	}

	err := prog.Validate(opts...)
	return err
}

func method5() error {
	// Overriding the constraints.
	const sample = `
#Min50: #Range & {
		min: >=50 
}

#Min50 & {
	"min": 10,
	"max": -10
}`
	ctx := cuecontext.New()
	bins := load.Instances([]string{"ranges.v2.cue"}, nil)
	schema := ctx.BuildInstance(bins[0])

	prog := ctx.CompileString(sample, cue.Scope(schema), cue.Filename("range.json"))
	opts := []cue.Option{
		cue.Attributes(true),
		cue.Definitions(true),
		cue.Hidden(true),
	}

	err := prog.Validate(opts...)
	fmt.Println(errors.Details(err, nil))
	return err

}
