package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	floatd := decimal.NewFromFloat(3.5)
	//decimal的float可以和decimal的int直接运算，真好
	floatd.Add(decimal.NewFromInt(9))

	fmt.Println(floatd)
	//3.5 这是一个典型错误，
	//这点， 和java的BigDecimal一样的，
	//必须要把add之后的结果赋值给原来的bigDecimal
	floatd = floatd.Add(decimal.NewFromInt(19))
	fmt.Println(floatd) //这才对
}
