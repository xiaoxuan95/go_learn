package method

// ShowCar 内嵌类的方法
func (self *car) ShowCar() (carInfo string) {
	carInfo = "汽车品牌：" + self.carBrand + "，汽车类型：" + self.carType
	return
}

// SetCarBrand set方法修改strcut的field
func (self *smallCar) SetCarBrand(newName string) {
	self.carBrand = newName
}

// GetCarBrand get方法获取struct的field
func (self *smallCar) GetCarBrand() (carBrand string) {
	carBrand = self.carBrand
	return
}

// ChangeType 更改汽车类型，手动、自动挡互换
func (self *smallCar) ChangeType() bool {
	if self.carType == "auto" {
		self.carType = "manual"
	} else if self.carType == "manual" {
		self.carType = "auto"
	} else {
		return false
	}
	return true
}
