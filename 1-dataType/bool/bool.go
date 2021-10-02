package bool

import "fmt"

func IsIntEqual(i, j int) (flag bool) {
	flag = i == j
	fmt.Printf("两个整数是否相等：%t\n", flag)
	return
}

func IsBothOver10(i, j int) (flag bool) {
	flag = (i > 10 && j > 10)
	fmt.Printf("两个整数是否都大于10：%t\n", flag)
	return
}

func IsEitherBMW(carname1, carname2 string) (flag bool) {
	flag = (carname1 == "BMW" || carname2 == "BMW")
	fmt.Printf("两部车中是否至少有一辆为BMW：%t\n", flag)
	return
}

func GetNot(b bool) (flag bool) {
	flag = !b
	fmt.Printf("布尔值%t取反结果为：%t\n", b, flag)
	return
}
