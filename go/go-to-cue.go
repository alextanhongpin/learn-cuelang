// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/gocode/gocodec"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	codec := gocodec.New(cuecontext.New(), nil)
	fmt.Println(codec.ExtractType(map[string]any{"name": "john", "age": 13}))
	fmt.Println(codec.ExtractType(`hello`))
	// We can extract the cue definition from a struct.
	fmt.Println(codec.ExtractType(User{}))
	fmt.Println(codec.ExtractType(&User{}))
	fmt.Println(codec.Decode(&User{}))
}
