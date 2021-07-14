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
