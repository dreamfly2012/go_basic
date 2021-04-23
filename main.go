package main

import (
	"fmt"
	"os"

	"github.com/dreamfly2012/go_basic/string"
	"github.com/google/go-cmp/cmp"
)

func main() {
	s := string.ReverseRunes("世界你好!")
	fmt.Println(s)
	diff := cmp.Diff("hello world", "hello go")
	fmt.Println(diff)

	useCAS := false
	if len(os.Args) > 1 && os.Args[1] == "cas" {
		useCAS = true
	}
	routineCnt := 4

	tryCnt := compareswap.cas._TEST_CNT / routineCnt
	var obj = &Obj{c: make(chan interface{}, cas._CHAN_SIZE+cas._GUARD_SIZE)}

	for idx := 0; idx < routineCnt; idx++ {
		go func(nameIdx int) {
			for tryIdx := 0; tryIdx < tryCnt; tryIdx++ {
				if useCAS {
					obj.writeMsgWithCASCheck(nameIdx, nil)
				} else {
					obj.writeMsg(nameIdx, nil)
				}
			}
		}(idx)
	}

	// fmt.Println(casObj.readLoop())
	fmt.Println(obj.readLoop())
	fmt.Println("quit.")
}
