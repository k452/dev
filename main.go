package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	for pos := 0; pos < 6; pos++ {
		for j := 0b0; j <= 0b11111; j++ { //5階差分
			//差分ベクトルにcを差し込む処理
			t := (j >> pos) & create2(5-pos)
			b := j & create2(pos)
			//rand.Seed(time.Now().UnixNano())
			in := 0b1
			output := (((t << 1) | in) << pos) | b
			fmt.Printf("%06b\n", output)
		}
		fmt.Println("--------------------")
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
