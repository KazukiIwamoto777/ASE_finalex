package main

import (
	"fmt"
	"strconv"
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

	if n[0] == n[1] || n[1] == n[2] || n[2] == n[0] {
		n = [...]int{0, 0, 0}
		return n
	}

	return n

}

func checkEatBite(n [3]int, m [3]int) (int, int) {
	var eat int
	var bite int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if n[i] == m[j] {
				if i == j {
					eat++
				} else {
					bite++
				}
			}
		}
	}
	return eat, bite
}

func stringtoInt(num string) int {
	//stringをintにする
	var num2 int
	num2, _ = strconv.Atoi(num)
	return num2
}

func main() {
	var n_m1 string
	print("input number: ")
	fmt.Scan(&n_m1)

	var n_m2 int
	n_m2 = stringtoInt(n_m1) //string をintにする関数

	var n [3]int
	n = numbertest(n_m2) //保持する数字のエラーを確かめる　配列に格納する
	fmt.Println(n)

	var n_est string
	print("input estimate number: ")
	fmt.Scan(&n_est)

	var n_est2 int
	n_est2 = stringtoInt(n_est)

	var m [3]int
	m = numbertest(n_est2) //推定する数字のエラーを確かめる　配列に格納する

	fmt.Println(m)

	eat, bite := checkEatBite(n, m)

	fmt.Print(eat)
	print("EAT ")
	fmt.Print(bite)
	println("BITE")
}
