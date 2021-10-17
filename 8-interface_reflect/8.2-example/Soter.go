package __2_example

type Soter interface {
	Len() int
	Compare(i, j int) bool
	Change(i, j int)
}

func Sort(data Soter) {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len()-i; j++ {
			if data.Compare(j, j+1) {
				data.Change(j, j+1)
			}
		}
	}
}

type IntList []int

func (self IntList) Len() int {
	return len(self)
}

func (self IntList) Compare(i, j int) bool {
	return self[i] > self[j]
}

func (self IntList) Change(i, j int) {
	self[i], self[j] = self[j], self[i]
}

type Day struct {
	index int
	name  string
}

type DayArray struct {
	data []*Day
}

func (self *DayArray) Len() int {
	return len(self.data)
}

func (self *DayArray) Compare(i, j int) bool {
	return self.data[i].index > self.data[j].index
}

func (self *DayArray) Change(i, j int) {
	self.data[i], self.data[j] = self.data[j], self.data[i]
}
