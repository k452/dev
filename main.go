package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	a := 0b0000000000000000000000000000000
	for i := 0; i < 32; i++ {
		t := (a >> i) & create2(i-1)
		b := a & create2(i)
		in := 0b1
		fmt.Printf("%032b\n", (((t<<1)|in)<<i)|b)
	}
	//for i := 0; i < 8; i++ {
	//	fmt.Printf("%b\n", (a>>(7-i))&1)
	//}

	fmt.Println("\n実行時間：", time.Since(start))
}

func create2(num int) int {
	txt := ""
	for j := 0; j < num; j++ {
		txt += "1"
	}
	res, _ := strconv.ParseInt(txt, 2, 32)
	return int(res)
}

func nSplit(msg string, splitlen int) []string {
	slc := []string{}
	for i := 0; i < len(msg); i += splitlen {
		if i+splitlen < len(msg) {
			slc = append(slc, msg[i:(i+splitlen)])
		} else {
			slc = append(slc, msg[i:])
		}
	}
	return slc
}
