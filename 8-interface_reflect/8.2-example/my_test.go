package __2_example

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	data1 := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	myList1 := IntList(data1)
	Sort(myList1)
	fmt.Println(data1)

	day1 := Day{0, "SUN"}
	day2 := Day{1, "MON"}
	day3 := Day{2, "TUE"}
	day4 := Day{3, "WED"}
	day5 := Day{4, "THU"}
	day6 := Day{5, "FRI"}
	day7 := Day{6, "SAT"}
	data2 := []*Day{&day4, &day6, &day5, &day2, &day7, &day1, &day3}
	myList2 := DayArray{data2}
	Sort(&myList2)
	for _, val := range data2 {
		fmt.Println(val)
	}
}
