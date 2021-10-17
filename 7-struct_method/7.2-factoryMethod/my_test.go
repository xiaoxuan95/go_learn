package __2_factoryMethod

import (
	"fmt"
	"testing"
)

func TestNewCar(t *testing.T) {
	car1 := NewCar("BMW", "auto")
	car2 := NewCar("LINKIN", "manual")
	fmt.Println(car1)
	fmt.Println(car2)
}
