package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	const sabun = 29

	output := 0b0
	pt := 0b0
	in := 0b111
	for pos := 0; pos <= 29; pos++ {
		l := 0b0
		r := 0b0

		if pos == 29 {
			l, r = split3(in, 0)
			output = (pt << 3) | l
			fmt.Println("pos:", pos)
			fmt.Printf("%032b\n", output)
		} else if pos == 0 {
			for sp := 0; sp < 3; sp++ {
				if sp == 0 {
					//posへlを投入
					output = (in << 29) | pt
				} else if sp == 1 {
					//posへl、pos+3へrを投入
					l, r = split3(in, sp)
					//output = (((l << 29) | pt) << 30) | pt
					output = ((l << 29) | pt)
					//fmt.Printf("aaa:%030b\n", output)
					left := (output >> 28) & 0b11                    //2bit
					right := output & 0b1111111111111111111111111111 //28bit
					output = (((left << 2) | r) << 28) | right
				} else if sp == 2 {
					//posへl、pos+2へrを投入
					l, r = split3(in, sp)
					output = ((l << 29) | pt)
					left := (output >> 29) & 0b111                    //1bit
					right := output & 0b11111111111111111111111111111 //29bit
					output = (((left << 2) | r) << 28) | right
				}
				fmt.Println("pos:", pos, "  sp:", sp)
				fmt.Printf("%032b\n", output)
			}
		} else {
			for sp := 0; sp < 3; sp++ {
				if sp == 0 {
					//posへlを投入
					left := (pt >> (29 - pos)) & create2(pos)
					right := pt & create2(29-pos)
					output = (((left << pos) | in) << (pos)) | right

					//fmt.Println("pos:", pos, "  sp:", sp)
					//fmt.Printf("%032b\n", output)
				} else if sp == 1 {
					//posへl、pos+3へrを投入
					l, r = split3(in, sp)
					left := (pt >> (29 - pos)) & create2(pos)
					right := pt & create2(29-pos)
					output = (((left << pos) | l) << pos) | right
					//fmt.Println("pos:", pos, "  sp:", sp)
					//fmt.Printf("%030b\n", output)

					if pos == 1 {
						output = (output << 2) | r
					} else {
						left2 := (output >> (pos - 1)) & create2(pos)
						right2 := output & create2(pos-1)
						//fmt.Printf("pos%dのoutput:%030b\n", pos, output)
						//fmt.Printf("pos%dのleft2:%0b\n", pos, left2)
						//fmt.Printf("pos%dのright2:%0b\n\n", pos, right2)
						output = (((left2 << 2) | r) << (pos - 1)) | right2

						//fmt.Println("pos:", pos, "  sp:", sp)
						//fmt.Printf("%032b\n", output)
					}
					//fmt.Println("pos:", pos, "  sp:", sp)
					//fmt.Printf("%032b\n", output)
				} else if sp == 2 {
					//posへl、pos+2へrを投入
					l, r = split3(in, sp)
					left := (pt >> (29 - pos)) & create2(pos)
					right := pt & create2(29-pos)
					output = (((left << pos) | l) << (pos)) | right
					//fmt.Println("pos:", pos, "  sp:", sp)
					//fmt.Printf("%030b\n", output)

					if pos == 1 {
						output = (output << 1) | r
					} else if pos == 2 {
						left2 := (output >> 1) & create2(3)
						right2 := output & create2(pos-1)
						//fmt.Printf("pos%dのoutput:%030b\n", pos, output)
						//fmt.Printf("pos%dのleft2:%0b\n", pos, left2)
						//fmt.Printf("pos%dのright2:%0b\n\n", pos, right2)
						output = (((left2 << 1) | r) << (pos - 1)) | right2
					} else {
						left2 := (output >> (pos - 1)) & create2(pos)
						right2 := output & create2(pos-1)
						//fmt.Printf("pos%dのoutput:%030b\n", pos, output)
						//fmt.Printf("pos%dのleft2:%0b\n", pos, left2)
						//fmt.Printf("pos%dのright2:%0b\n\n", pos, right2)
						output = (((left2 << 1) | r) << (pos - 1)) | right2
					}
					//fmt.Println("pos:", pos, "  sp:", sp)
					//fmt.Printf("%032b\n", output)
				}
				fmt.Println("pos:", pos, "  sp:", sp)
				fmt.Printf("%032b\n", output)
			}
		}
	}

	/*
			t := (in >> pos) & create2(28-pos)
			b := in & create2(pos)
			fmt.Printf("%b\n", t)
			fmt.Printf("%b\n", b)


		if pos == 0 {
			output = ((t << 3) | in) << pos
		} else if pos == 28 {
			output = (in << 28) | b
		} else {
			output = (((t << 3) | in) << pos) | b
		}
		fmt.Printf("%032b\n", output)
	*/

	//pos := 28

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
