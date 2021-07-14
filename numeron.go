package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func numbertest(n_m int) [3]int {
	var n [3]int
	if n_m < 12 || n_m > 987 {
		//エラー
		//[0,0,0]の配列を返す
		n = [...]int{0, 0, 0}
		return n
	}

	n[0] = n_m / 100
	n_m = n_m - 100*n[0]
	n[1] = n_m / 10
	n[2] = n_m - 10*n[1]

	return n

}

func main() {
	var n_m1 int
	print("input number: ")
	fmt.Scan(&n_m1)

	var n [3]int
	n = numbertest(n_m1)

	fmt.Println(n)

}
