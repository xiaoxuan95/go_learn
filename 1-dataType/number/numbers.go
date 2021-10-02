package number

import "fmt"
import "github.com/shopspring/decimal"

func AddInt(i, j int) (sum int) {
	sum = i + j
	fmt.Printf("两整数之和为：%d\n", sum)
	return
}

func FloatDecimal(i, j float64) (sum float64) {
	dcmlFloat1, dcmlFloat2 := decimal.NewFromFloat(i), decimal.NewFromFloat(j)
	dcmlSum := dcmlFloat1.Add(dcmlFloat2)
	sum = decimalTofloat(dcmlSum)
	return
}

func decimalTofloat(d decimal.Decimal) float64 {
	f, exact := d.Float64()
	if !exact {
		return f
	}
	return 0
}

func decimalToInt(d decimal.Decimal) int64 {
	return d.IntPart()
}
