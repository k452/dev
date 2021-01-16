package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, _ := os.Open("./16kai.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	text := []string{"1111", "1111", "1111", "1111"}

	for scanner.Scan() {
		//ここから開始
		tmp := strings.Split(scanner.Text(), ",")
		org := []string{"0", "0", "0", "0", "0", "0", "0", "0"}
		posA := strings.Split(tmp[0], "")
		posC := strings.Split(tmp[1], "")

		for i := 0b0; i <= 0b1111111111111111; i++ { //16階差分
			sabun := nSplit(fmt.Sprintf("%016b", i), 4)
			for i, v := range posA {
				v, _ := strconv.Atoi(v)
				org[v] = sabun[i]
			}
			for i, v := range posC {
				v, _ := strconv.Atoi(v)
				org[v] = text[i]
			}
			pt, _ := strconv.ParseInt(strings.Join(org, ""), 2, 32)
			pt ^= 0
			//fmt.Printf("%032b\n", pt)
		}

	}

	fmt.Println("\n実行時間：", time.Since(start))
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
