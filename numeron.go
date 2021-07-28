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

func errorcheck(num [3]int) bool {
	if num[0] == 0 && num[1] == 0 && num[2] == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	var p1string string
	var p1int int
	var p1 [3]int

	for {
		print("Player1 : ")
		fmt.Scan(&p1string)

		p1int = stringtoInt(p1string)
		p1 = numbertest(p1int) //保持する数字のエラーを確かめる　配列に格納する

		if errorcheck(p1) == true {
			break
		} else {
			fmt.Println("Please enter the number again :")
		}
	}
	var p2string string
	var p2int int
	var p2 [3]int

	for {
		print("Player2 : ")
		fmt.Scan(&p2string)

		p2int = stringtoInt(p2string)
		p2 = numbertest(p2int) //保持する数字のエラーを確かめる　配列に格納する

		if errorcheck(p2) == true {
			break
		} else {
			fmt.Println("Please enter the number again")
		}
	}

	for {
		//p1 turn
		var p1est_string string
		var p1est_int int
		var p1est [3]int

		fmt.Println("Player1のターンです")
		for {
			fmt.Print("Estimate number : ")
			fmt.Scan(&p1est_string)

			p1est_int = stringtoInt(p1est_string)
			p1est = numbertest(p1est_int) //保持する数字のエラーを確かめる　配列に格納する

			if errorcheck(p1est) == true {
				break
			} else {
				fmt.Println("Please enter the number again")
			}
		}
		//eat,biteの表示

		//p2 turn

		var p2est_string string
		var p2est_int int
		var p2est [3]int
		fmt.Println("Player2のターンです")
		for {
			fmt.Print("Estimate number : ")
			fmt.Scan(&p2est_string)

			p2est_int = stringtoInt(p2est_string)
			p2est = numbertest(p2est_int) //保持する数字のエラーを確かめる　配列に格納する

			if errorcheck(p2est) == true {
				break
			} else {
				fmt.Println("Please enter the number again")
			}
		}
	}
	/*
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
	*/
}
