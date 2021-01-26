package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	output := 0b0
	for pos := 0; pos < 6; pos++ {
		for i := 0; i <= 0b11111; i++ {
			//i := 0b0
			t := (i >> pos) & create2(8-pos)
			b := i & create2(pos)
			in := 0b111

			if pos == 0 {
				output = ((t << 3) | in) << pos
			} else {
				output = (((t << 3) | in) << pos) | b
			}
			fmt.Printf("%08b\n", output)
		}
	}

	fmt.Println("\n実行時間：", time.Since(start))
}

func split3(in int, pos int) (int, int) {
	l := 0b0
	r := 0b0

	if pos == 0 {
		l = in
	} else if pos == 1 {
		l = (in >> 2) & create2(0b1)
		r = in & 0b11
	} else if pos == 2 {
		l = (in >> 1) & create2(0b11)
		r = in & 0b1
	}

	return l, r
}

func create2(num int) int {
	res := 0b0
	switch num {
	case 0:
		res = 0b0
		break
	case 1:
		res = 0b1
		break
	case 2:
		res = 0b11
		break
	case 3:
		res = 0b111
		break
	case 4:
		res = 0b1111
		break
	case 5:
		res = 0b11111
		break
	case 6:
		res = 0b111111
		break
	case 7:
		res = 0b1111111
		break
	case 8:
		res = 0b11111111
		break
	case 9:
		res = 0b111111111
		break
	case 10:
		res = 0b1111111111
		break
	case 11:
		res = 0b11111111111
		break
	case 12:
		res = 0b111111111111
		break
	case 13:
		res = 0b1111111111111
		break
	case 14:
		res = 0b11111111111111
		break
	case 15:
		res = 0b111111111111111
		break
	case 16:
		res = 0b1111111111111111
		break
	case 17:
		res = 0b11111111111111111
		break
	case 18:
		res = 0b111111111111111111
		break
	case 19:
		res = 0b1111111111111111111
		break
	case 20:
		res = 0b11111111111111111111
		break
	case 21:
		res = 0b111111111111111111111
		break
	case 22:
		res = 0b1111111111111111111111
		break
	case 23:
		res = 0b11111111111111111111111
		break
	case 24:
		res = 0b111111111111111111111111
		break
	case 25:
		res = 0b1111111111111111111111111
		break
	case 26:
		res = 0b11111111111111111111111111
		break
	case 27:
		res = 0b111111111111111111111111111
		break
	case 28:
		res = 0b1111111111111111111111111111
		break
	case 29:
		res = 0b11111111111111111111111111111
		break
	case 30:
		res = 0b111111111111111111111111111111
		break
	case 31:
		res = 0b1111111111111111111111111111111
		break
	case 32:
		res = 0b11111111111111111111111111111111
		break
	default:
		panic("範囲外")
	}
	return res
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
