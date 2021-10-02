package control

import "fmt"

var Num int

func SwtichTest(x int) int {
	val := x % 2
	switch val {
	case 0:
		fmt.Printf("输入为偶数:%d，两倍值为：", x)
		Num = x * 2
	case 1:
		fmt.Printf("输入为奇数:%d，三倍值为：", x)
		Num = x * 3
	}
	return Num
}

func SwitchTest2(num int) {
	switch num1 := num; {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}
