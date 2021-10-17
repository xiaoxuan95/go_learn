package __3_inherit

import (
	"fmt"
	"testing"
)

func TestNewCar(t *testing.T) {
	myCar := NewCar("small", "EES015", "Chevrolet", "auto")
	fmt.Println(myCar)
}
