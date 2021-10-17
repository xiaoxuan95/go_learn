package __4_method

type car struct {
	carBrand string
	carType  string
}

type smallCar struct {
	int
	license string
	car
}

func NewCar(class string, license string, carBrand string, carType string) *smallCar {
	if class == "small" {
		return &smallCar{1, license, car{carBrand: carBrand, carType: carType}}
	}
	return nil
}
