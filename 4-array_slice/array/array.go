package array

//FirstArray
//声名一个数组
func FirstArray1() [5]string {
	var firstArray [5]string
	firstArray[0] = "val_1"
	firstArray[1] = "val_2"
	firstArray[2] = "vak_3"
	return firstArray
}

//FirstArray2
//声名数组的另外一种方式
func FirstArray2() [5]int {
	firstArray := [5]int{0, 1, 2, 3}
	return firstArray
}

// FirstArray3
//声名数组的最后一种方式
func FirstArray3() [5]float64 {
	firstArray := [5]float64{3: 1.23, 4: 3.14}
	return firstArray
}
