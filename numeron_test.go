package main

import "testing"

func Testnumeron01(t *testing.T) {
	str := numbertest(123)
	expect := [3]int{1, 2, 3}
	if str != expect {
		t.Error("Test01 is failed")
	}
}

func Testnumeron02(t *testing.T) {
	str := numbertest(1234)
	expect := [3]int{0, 0, 0}
	if str != expect {
		t.Error("Test02 is failed")
	}
}

func Testnumeron03(t *testing.T) {
	str := numbertest(222)
	expect := [3]int{0, 0, 0}
	if str != expect {
		t.Error("Test03 is failed")
	}
}

func Testnumeron04(t *testing.T) {
	eat, bite := checkEatBite([3]int{1, 7, 6}, [3]int{1, 3, 5})
	if eat != 1 && bite != 1 {
		t.Error("Test04 is failed")
	}
}
