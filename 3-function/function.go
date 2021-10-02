package function

import "log"

//SearchMinInt
//函数传递可变参数，传递slice切片
//查找最小的整数
func SearchMinInt(s ...int) (min int) {
	min = s[0]
	for _, val := range s {
		if val < min {
			min = val
		}
	}
	return
}

//DoubleInt
//函数通过defer跟踪入参、返回值
//整数翻一番
func DoubleInt(intPtr *int) (doubleVal int) {
	inputVal := *intPtr
	*intPtr = *intPtr * 2
	doubleVal = *intPtr
	defer func() {
		log.Printf("DoubleInt(%d) = %d", inputVal, *intPtr)
	}()
	return
}

//Finb1
//递归实现斐波那契
func Finb1(i int) (res int) {
	if i <= 1 {
		res = 1
	} else {
		res = Finb1(i-1) + Finb1(i-2)
	}
	return
}

//Finb2
//闭包实现斐波那契
func Finb2() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
