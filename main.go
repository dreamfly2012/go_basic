package main

import (
	"fmt"

	"github.com/dreamfly2012/go_basic/string"
	"github.com/dreamfly2012/go_basic/time"
	"github.com/google/go-cmp/cmp"
)

func main() {
	s := string.ReverseRunes("世界你好!")
	fmt.Println(s)
	diff := cmp.Diff("hello world", "hello go")
	fmt.Println(diff)
	time := time.FormatTime("2021-10-12 12:12:12")
	fmt.Println(time)
}
