package __3_emptyInterface

import "fmt"

func Hello() {
	fmt.Println("welcome")
}

type Person struct {
	name string
	age  int
}

type Any interface{}

type AnyList struct {
	anylist []Any
}

func (self *AnyList) Add(index int, element Any) {
	self.anylist[index] = element
}

func (self *AnyList) Get(index int) Any {
	return self.anylist[index]
}
