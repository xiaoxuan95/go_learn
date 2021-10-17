package __1_defineMethod

type FirstStruct struct {
	name string
	age  int
}

// DefineStruct1 结构体的定义方法1
func DefineStruct1() FirstStruct {
	var myStr FirstStruct
	myStr.name = "xiaogao"
	return myStr
}

// DefineStruct2 结构体的定义方法2，new关键字
func DefineStruct2() *FirstStruct {
	myStr := new(FirstStruct)
	myStr.name = "gzx"
	return myStr
}

// DefineStruct3 结构体的定义方法3，简短式
func DefineStruct3() *FirstStruct {
	myStr := &FirstStruct{name: "gaozhixuan", age: 25}
	return myStr
}
