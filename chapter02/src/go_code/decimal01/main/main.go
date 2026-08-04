package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	var num float32 = 120.34
	var divider int64 = 5
	//fmt.Println(num/divider)// Invalid operation: num/divider (mismatched types float32 and int64)
	//如果不用decimal类，float32和int64是不能直接做除法的

	numd := decimal.NewFromFloat32(num)
	dividerd := decimal.NewFromInt(divider)
	fmt.Println(numd.DivRound(dividerd, 4))
	//用了decimal类，decimal类本身并不区分是来自于float32还是int64，反正只要转成了decimal类，就可以相互运算
}
