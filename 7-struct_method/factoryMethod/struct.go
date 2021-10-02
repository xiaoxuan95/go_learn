package factoryMethod

type car struct {
	carBrand string
	carType  string
}

func NewCar(carBrand string, carType string) *car {
	if carType != "auto" && carType != "manual" {
		return nil
	}
	if carBrand != "BMW" && carBrand != "AUDI" && carBrand != "BENZ" {
		return nil
	}
	return &car{carBrand: carBrand, carType: carType}
}
