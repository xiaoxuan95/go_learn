package string

import (
	"fmt"
	"strconv"
	"strings"
)

func StringsTest() {
	// 前、后缀
	strings1 := "start xxxx end"
	myFlag := strings.HasPrefix(strings1, "start") &&
		strings.HasSuffix(strings1, "end")
	fmt.Printf("字符串前后缀是否合法：%t\n", myFlag)
	// 是否包含某字符，找出其位置
	fmt.Println(strings.Contains("abcdedf", "abc"))
	fmt.Println(strings.Index("a123a123", "123"),
		strings.LastIndex("123a123", "123"))
	// 批量替换某字符
	string1 := "old1old2old3"
	string2 := strings.Replace(string1, "old", "new", -1)
	string3 := strings.Replace(string1, "old", "new", 2)
	fmt.Println(string2, string3)
	// 包含的个数
	fmt.Println(strings.Count("123451", "1"))
	// 重复
	fmt.Println(strings.Repeat("abc", 3))
	// 字母大小写
	string4 := "aBcD"
	fmt.Println(strings.ToLower(string4), strings.ToUpper(string4))
	// 去除空字符
	fmt.Println(strings.TrimSpace(" 修剪后的字符串\n		"))
	// 按指定字符切割,合并
	string5 := "1995-10-11"
	fmt.Println(strings.Split(string5, "-"))
	slice1 := strings.Split(string5, "-")
	for _, x := range slice1 {
		fmt.Println(x)
	}
	fmt.Println(strings.Join(slice1, ":"))
	// 基础类型互相转换
	int1 := 123
	float1 := 3.141531264754634
	bool1 := true
	fmt.Println(strconv.FormatInt(int64(int1), 10))
	fmt.Println(strconv.FormatFloat(float1, 'f', -1, 64))
	fmt.Println(strconv.FormatBool(bool1))
	str1 := "123"
	str2 := "3.14"
	str3 := "false"
	fmt.Println(strconv.Atoi(str1))
	fmt.Println(strconv.ParseFloat(str2, 64))
	fmt.Println(strconv.ParseBool(str3))
}
