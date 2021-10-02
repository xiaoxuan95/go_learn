package function

import (
	"fmt"
	"testing"
)

func TestSearchMinInt(t *testing.T) {
	intSlice := []int{6, 3, 5, 1, 2, 0, 8, 9}[:]
	fmt.Println(SearchMinInt(intSlice...))
}

func TestDoubleInt(t *testing.T) {
	numInt := 1
	fmt.Println(DoubleInt(&numInt))
	fmt.Println(DoubleInt(&numInt))
	fmt.Println(DoubleInt(&numInt))
}

func TestFinb1(t *testing.T) {
	for x := 1; x <= 10; x++ {
		fmt.Println(Finb1(x))
	}
}

func TestFinb2(t *testing.T) {
	fin := Finb2()
	num := 1
	for {
		fmt.Println(fin())
		if num > 10 {
			break
		}
		num = num + 1
	}
}
