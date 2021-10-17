package __4_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	var x int
	x = 123
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("val:", reflect.ValueOf(x))
}

func TestHello2(t *testing.T) {
	type car struct {
		Carid   int
		Carname string
	}

	car1 := car{1, "my_car1"}
	s := reflect.ValueOf(&car1).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println(f.Type(), f.Interface())
	}

}
