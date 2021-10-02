package slice

import (
	"fmt"
	"testing"
)

func TestFirstSlice1(t *testing.T) {
	fmt.Println("我声明的第一个slice：", FirstSlice1())
	for index, val := range FirstSlice1() {
		fmt.Printf("index:%d, val:%d\n", index, val)
	}
}

func TestFirstSlice2(t *testing.T) {
	FirstSlice2()
}

func TestFirstSlice3(t *testing.T) {
	slice1 := FirstSlice3()
	fmt.Println("我声明的第一个slice：", slice1)
}

func TestFirstSlice4(t *testing.T) {
	slice1 := FirstSlice4()
	fmt.Printf("make创建的切片,：%v\n长度为：%d, 容量为：%d\n", slice1, len(slice1), cap(slice1))
	slice1[1], slice1[2] = 1, 2
	fmt.Println("重新赋值后的切片：", slice1)
	slice1 = slice1[:cap(slice1)]
	fmt.Println("扩容后的切片：", slice1)
}
