package slice

import "fmt"

// FirstSlice1
//声名slice的一种方式
func FirstSlice1() []int {
	var firstSlice []int
	firstSlice = []int{0: 1, 2: 3}
	return firstSlice
}

// FirstSlice2
//根据array初始化slice
func FirstSlice2() {
	myArray := [5]int{0: 0, 1: 1, 3: 3}
	fmt.Println("原数组为：", myArray)
	mySlice1 := myArray[1:4]
	fmt.Printf("index取1~3，生成的slice：%v\n长度为：%d，容量为：%d\n", mySlice1, len(mySlice1), cap(mySlice1))
	mySlice2 := myArray[:]
	fmt.Printf("全部元素的slice：%v\n长度为：%d，容量为：%d\n", mySlice2, len(mySlice2), cap(mySlice2))
	mySlice3 := myArray[3:]
	fmt.Printf("index取3~4，生成的slice：%v\n长度为：%d，容量为：%d\n", mySlice3, len(mySlice3), cap(mySlice3))
}

//FirstSlice3
//直接初始化slice
func FirstSlice3() []string {
	mySlice := []string{"val1", "val2", "val3"}
	return mySlice
}

//FirstSlice4
//make初始化slice
func FirstSlice4() []int {
	mySlice := make([]int, 3, 5)
	return mySlice
}
