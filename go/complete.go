package main

import (
	"log"
	"strings"

	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cuego"
)

func main() {
	type Sum struct {
		A int `cue:"C-B" json:",omitempty"`
		B int `cue:"C-A" json:",omitempty"`
		C int `cue:"A+B" json:",omitempty"`
	}

	a := Sum{A: 1, B: 5}
	err := cuego.Complete(&a)
	log.Printf("%+v (err: %s)\n", a, errMsg(err))

	b := Sum{A: 2, B: 8}
	err = cuego.Complete(&b)
	log.Printf("%+v (err: %s)\n", b, errMsg(err))

	c := Sum{A: 2, B: 3, C: 8}
	err = cuego.Complete(&c)
	log.Printf("%+v (err: %s)\n", c, errMsg(err))
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
