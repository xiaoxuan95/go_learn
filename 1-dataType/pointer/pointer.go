package pointer

import "fmt"

func PrtFuc(x, y *int) int {
	*x = 2 * *x
	*y = *y + 1
	fmt.Printf("修改后的x值:%d\n", *x)
	fmt.Printf("修改后的y值:%d\n", *y)
	return *x + *y
}
