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

func Testnumeron05(t *testing.T) {
	eat, bite := checkEatBite([3]int{0, 1, 2}, [3]int{0, 3, 4})
	if eat != 1 && bite != 1 {
		t.Error("Test05 is failed")
	}
}

func Testnumeron06(t *testing.T) {
	check := errorcheck([3]int{0, 1, 2})
	expect := true
	if check != expect {
		t.Error("Test06 is failed")
	}
}

func Testnumeron07(t *testing.T) {
	check := errorcheck([3]int{0, 0, 0})
	expect := false
	if check != expect {
		t.Error("Test07 is failed")
	}
}
