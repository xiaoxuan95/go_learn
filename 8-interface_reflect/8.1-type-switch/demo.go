package type_switch

type Factory interface {
	Build() map[string]string
}

type Benz struct {
	carType string
}

type Audi struct {
	carType  string
	maxSpeed string
}

func (self *Benz) Build() map[string]string {
	infoMap := make(map[string]string)
	infoMap["carType"] = self.carType
	return infoMap
}

func (self *Audi) Build() map[string]string {
	infoMap := make(map[string]string)
	infoMap["carType"] = self.carType
	infoMap["maxSpeed"] = self.maxSpeed
	return infoMap
}

func TypeSwitch(factory Factory) (typeInfo string) {
	switch factory.(type) {
	case *Benz:
		typeInfo = "该实例的类型为：Benz"
	case *Audi:
		typeInfo = "该实例的类型为：Audi"
	default:
		typeInfo = "未知类型"
	}
	return
}
