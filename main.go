package main

import (
	"fmt"

	"github.com/dreamfly2012/go_basic/mapinfo"
	"github.com/dreamfly2012/go_basic/string"
	"github.com/google/go-cmp/cmp"
)

func main() {
	s := string.ReverseRunes("世界你好!")
	fmt.Println(s)
	diff := cmp.Diff("Hello World", "hello go")
	fmt.Println(diff)
	m := mapinfo.GenerateMap()
	fmt.Println(m)

	slices := string.Slice(3)
	slices[0] = 0
	slices[1] = 1
	slices[2] = 2
	fmt.Println(slices)
}
