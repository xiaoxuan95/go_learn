package method

import (
	"fmt"
	"testing"
)

func TestSmallCar_ChangeType(t *testing.T) {
	myCar := NewCar("small", "EES015", "Chevrolet", "auto")
	myCar.ChangeType()
	fmt.Println(myCar)

}

func TestSmallCar_GetCarBrand(t *testing.T) {
	myCar := NewCar("small", "EES015", "Chevrolet", "auto")
	myCar.SetCarBrand("BMW")
	fmt.Println(myCar.GetCarBrand())
}

func TestCar_ShowCar(t *testing.T) {
	myCar := NewCar("small", "EES015", "Audi", "auto")
	fmt.Println(myCar.ShowCar())
}
