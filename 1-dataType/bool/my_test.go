package bool

import (
	"testing"
)

func TestIsIntEqual(t *testing.T) {
	i := 3
	j := 2
	IsIntEqual(i, j)
}

func TestIsBothOver10(t *testing.T) {
	i := 100
	j := 11
	IsBothOver10(i, j)
}

func TestIsEitherBMW(t *testing.T) {
	carname1 := "BMW"
	carname2 := "ABCX"
	IsEitherBMW(carname1, carname2)
}

func TestGetNot(t *testing.T) {
	b1, b2 := true, false
	GetNot(b1)
	GetNot(b2)
}
