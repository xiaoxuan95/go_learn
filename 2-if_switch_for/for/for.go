package control

import (
	"fmt"
	"strings"
)

// 基于计数器的循环

// TestFor1
//写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字， 但打印一次 "Fizz"。
//遇到 5 的倍数时，打印 Buzz 而不是相应的数字。
//对于同时为 3 和 5 的倍数的数，打印 FizzBuzz（提示：使用 switch 语句）。
func TestFor1() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("循环到了数字：%d\n", i)
		switch {
		case i%3 == 0 && i%5 != 0:
			fmt.Println("Fizz")
		case i%3 != 0 && i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		}
	}
}

// TestFor2
//使用 * 符号打印宽为 20，高为 10 的矩形。
func TestFor2() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 基于条件判断

func TestFor3() {
	i := 5
	for i >= 0 {
		fmt.Printf("The variable i is now: %d\n", i)
		i = i - 1
	}
}

//无限循环

func TestFor4() {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Printf("The variable i is now: %d\n", i)
		i = i + 1

	}
}

// for-range

func TestFor5() {
	string := "go is a beautiful language"
	slice := strings.Split(string, " ")
	for index, val := range slice {
		fmt.Println(index, val)
	}
}
