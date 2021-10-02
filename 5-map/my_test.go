package __map

import (
	"fmt"
	"testing"
)

func TestInitMap1(t *testing.T) {
	myMap := InitMap1()
	for key, value := range myMap {
		fmt.Printf("key值：%v, 对应的val值：%v\n", key, value)
	}
}

func TestInitMap2(t *testing.T) {
	myMap := InitMap2()
	fmt.Println("map中存的值分别为：", myMap["key1"], myMap["key2"], myMap["key-999"])
}

func TestIfKeyExist(t *testing.T) {
	myMap := make(map[string]string)
	key1, key2 := "my-key1", "my-key2"
	myMap[key1], myMap[key2] = "val1", "val2"
	fmt.Printf("map：%v，是否存在值为%v的key？%v\n", myMap, key1, IfKeyExist(myMap, key1))
	fmt.Printf("map：%v，是否存在值为%v的key？%v\n", myMap, key2, IfKeyExist(myMap, key2))
	delete(myMap, key1)
	fmt.Printf("map：%v，是否存在值为%v的key？%v\n", myMap, key1, IfKeyExist(myMap, key1))
	fmt.Printf("map：%v，是否存在值为%v的key？%v\n", myMap, key2, IfKeyExist(myMap, key2))

}
