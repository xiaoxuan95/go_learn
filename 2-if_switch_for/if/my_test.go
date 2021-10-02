package control

import (
	"fmt"
	"testing"
)

func TestIfTest(t *testing.T) {
	fmt.Println(IfTest(false))
}

func TestStrPlus(t *testing.T) {
	StrPlus("123")
	StrPlus("12x")
}
