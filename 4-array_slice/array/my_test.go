package array

import (
	"fmt"
	"testing"
)

func TestFirstArray1(t *testing.T) {
	fmt.Println("我声名的第一个字符串数组：", FirstArray1())
	for index, val := range FirstArray1() {
		fmt.Printf("index：%d，val：%s\n", index, val)
	}
}

func TestFirstArray2(t *testing.T) {
	fmt.Println("我声名的第一个整数数组：", FirstArray2())
	for index, val := range FirstArray2() {
		fmt.Printf("index：%d，val：%d\n", index, val)
	}
}

func TestFirstArray3(t *testing.T) {
	fmt.Println("我声名的第一个浮点数数组：", FirstArray3())
	for index, val := range FirstArray3() {
		fmt.Printf("index：%d，val：%g\n", index, val)
	}
}
