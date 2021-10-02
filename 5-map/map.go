package __map

//InitMap1
//map的初始化方法1
func InitMap1() (firstMap map[int]string) {
	firstMap = map[int]string{0: "val1", 1: "val2"}
	return
}

//InitMap2
//map的初始化方法2
func InitMap2() map[string][]int {
	firstMap := make(map[string][]int)
	firstMap["key1"] = []int{1, 2, 3}
	firstMap["key2"] = []int{4, 5, 6}
	return firstMap
}

// IfKeyExist 检测key是否存在
func IfKeyExist(testMap map[string]string, key string) bool {
	_, ok := testMap[key]
	return ok
}
