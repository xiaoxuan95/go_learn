package number

import (
	"fmt"
	"testing"
)

func TestAddInt(t *testing.T) {
	i := 3
	j := 4
	AddInt(i, j)
}

func TestFloatDecimal(t *testing.T) {
	i := 0.3
	j := 0.6
	fmt.Println("不使用decimal时，0.3+0.6的值：", i+j)
	fmt.Println("使用decimal时，0.3+0.6的值：", FloatDecimal(i, j))
}
