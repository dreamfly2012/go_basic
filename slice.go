package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 1
	s[1] = 10
	s[2] = 4

	s = append(s, 8)
	fmt.Println(s)
}
