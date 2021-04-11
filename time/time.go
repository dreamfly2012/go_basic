package main

import (
	"fmt"
	"time"
)

func main() {
	mytime, _ := time.Parse("2016-01-23 23:23:23", "2020-04-10 18:15:10")
	fmt.Println(mytime)
	now := time.Now()
	fmt.Println(now)
	format_now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(format_now)
	fmt.Println("vim-go")
}
