package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	//arr := []string{}
	/*
		x, _ := strconv.ParseInt("0001", 2, 32)
		for j := 0; j < 8; j++ {
			for i := 0b0; i <= 0b11; i++ {

					if strings.Index(fmt.Sprintf("%032b", i), "11111111111110") == 0 {
						arr = append(arr, fmt.Sprintf("%032b", i))
					}

				fmt.Printf("%08b\n", RotateL8(i, j))
			}
		}
	*/
	m := map[string]int{}
	o := []int{}
	for i := 0; i < 8; i++ {
		//O = map[string]int{"0":0, "1":1, "2":2, "3":3, "4":4, "5":5, "6":6, "7":7}
		m[""+fmt.Sprintf("%d", i)+""] = i
		o = append(o, i)
	}

	com := combination(o, 7)
	for _, v := range com {
		m = map[string]int{}
		for i := 0; i < 8; i++ {
			//O = map[string]int{"0":0, "1":1, "2":2, "3":3, "4":4, "5":5, "6":6, "7":7}
			m[""+fmt.Sprintf("%d", i)+""] = i
			o = append(o, i)
		}
		for _, w := range v {
			delete(m, fmt.Sprintf("%d", w))
		}
		fmt.Println(v, m)
	}

	//A := []int{1, 3, 5, 7}
	//NA := O
	//fmt.Println(combination(O, 4)[0])

	//file, _ = os.Create("./sample.csv")
	//c := []string{"0", "1", "0", "1"}
	/*
		for j := 0; j < 1; j++ {
			for k := 0b0; k <= 0b1111; k++ {
				arr := strings.Split(fmt.Sprintf("%04b", k), "")
				origin := []string{"0", "0", "0", "0", "0", "0", "0", "0"}
				for i, v := range A {
					origin[v] = arr[i]
				}
				for i, v := range NA {
					origin[v] = c[i]
				}
				//fmt.Println(strings.Join(origin, ""))
			}
		}
	*/
	fmt.Println("\n実行時間：", time.Since(start))
}

//以下2つが任意の範囲で乱数を生成して配列に格納する関数
func k(m map[int]bool) []int {
	i := 0
	result := make([]int, len(m))
	for key := range m {
		result[i] = key
		i++
	}
	return result
}

func random(min int, max int, num int) []int {
	numRange := max - min + 1
	selected := make(map[int]bool)
	rand.Seed(time.Now().UnixNano())
	for counter := 0; counter < num; {
		n := rand.Intn(numRange) + min
		if selected[n] == false {
			selected[n] = true
			counter++
		}
	}
	keys := k(selected)
	sort.Sort(sort.IntSlice(keys)) // ソートしない場合コメントアウト
	return keys
}

//RotateL8 8bitの数値を対象とした左シフト
func RotateL8(a int, i int) int {
	return ((a<<i)&0xff ^ (a >> (8 - i)))
}

//RotateL32 32bitの数値を対象とした左シフト
func RotateL32(a int, i int) int {
	return ((a<<i)&0xffffffff ^ (a >> (32 - i)))
}

//組み合わせ
func combination(ar []int, n int) (result [][]int) {
	if n <= 0 || len(ar) < n {
		return
	}
	if n == 1 {
		for _, a := range ar {
			result = append(result, []int{a})
		}
	} else if len(ar) == n {
		result = append(result, ar)
	} else {
		for _, a := range combination(ar[1:], n-1) {
			result = append(result, append([]int{ar[0]}, a...))
		}
		result = append(result, combination(ar[1:], n)...)
	}
	return
}
