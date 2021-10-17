package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

type User struct {
	Name string
	Age  int
}

func main() {
	ctx := cuecontext.New()
	p := ctx.CompileString("a: 2")
	fmt.Println(p)

	s, _ := p.Fields()
	for s.Next() {
		sel := s.Selector()
		fmt.Println("sel:", sel)

		path := cue.MakePath(sel)
		fmt.Println("path:", path)

		v := p.LookupPath(path)
		if v.Exists() {
			fmt.Println("lookup:", v)
		}
	}

	switch p.IncompleteKind() {
	case cue.StructKind:
		fmt.Println("is struct")
	case cue.ListKind:
		fmt.Println("is list")
	}
}
