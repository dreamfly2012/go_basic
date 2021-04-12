package main

import (
	"fmt"

	"github.com/dreamfly2012/go_basic/string"
	"github.com/google/go-cmp/cmp"
)

func main() {
	s := string.ReverseRunes("世界你好!")
	fmt.Println(s)
	diff := cmp.Diff("hello world", "hello go")
	fmt.Println(diff)
}
