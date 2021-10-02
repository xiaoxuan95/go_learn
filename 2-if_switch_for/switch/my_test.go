package control

import (
	"fmt"
	"testing"
)

func TestSwtichTest(t *testing.T) {
	fmt.Println(SwtichTest(4))
	fmt.Println(SwtichTest(1))
}

func TestSwitchTest2(t *testing.T) {
	SwitchTest2(-1)
	SwitchTest2(5)
	SwitchTest2(15)
}
