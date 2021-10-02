package defineMethod

import (
	"fmt"
	"testing"
)

func TestDefineStruct1(t *testing.T) {
	myStr := DefineStruct1()
	fmt.Println(myStr)
	fmt.Println(myStr.name)
}

func TestDefineStruct2(t *testing.T) {
	myStr := DefineStruct2()
	fmt.Println(myStr)
	fmt.Println("该实例的地址：", &myStr)
	fmt.Println(myStr.name)
}

func TestDefineStruct3(t *testing.T) {
	myStr := DefineStruct3()
	fmt.Println(myStr)
	fmt.Println("该实例的地址：", &myStr)
}
