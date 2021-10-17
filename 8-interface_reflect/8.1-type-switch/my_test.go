package type_switch

import (
	"fmt"
	"testing"
)

func TestTypeSwitch(t *testing.T) {
	var fac1 Factory
	var fac2 Factory

	car1 := new(Benz)
	car1.carType = "manual"

	car2 := new(Audi)
	car2.carType = "auto"
	car2.maxSpeed = "300"

	fac1 = car1
	fac2 = car2

	fmt.Println(TypeSwitch(fac1))
	fmt.Println(fac1.Build())
	fmt.Println(TypeSwitch(fac2))
	fmt.Println(fac2.Build())
}
