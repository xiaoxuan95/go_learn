package pointer

import (
	"fmt"
	"testing"
)

var (
	num1 int
	num2 int
)

func TestPrtFuc(t *testing.T) {
	num1 = 1
	num2 = 2
	n := 1
	for n <= 3 {
		fmt.Println(PrtFuc(&num1, &num2))
		n = n + 1
	}
}
