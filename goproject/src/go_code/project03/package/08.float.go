package main

import "fmt"

func main() {
	//var price float32 = 13.89
	//fmt.Println(price)

	//	单精度（java里的float） vs 双精度（java里的double）
	//	golang的float32和float64

	//go的普通业务计算一般优先使用float64，float32反而只在明确需要节省内存时使用（不建议用）
	//goland格式化代码

	//Go 不允许同一个 utils 中有多个 main 函数。

	//在同一个目录下有多个含func main()的.go文件时，不能一起编译，只能用go run 具体文件 来单独运行某一个

	//最像java的bigDecimal的库是decimal
	//price := decimal.RequireFromString("136.02")
	//quantity := decimal.NewFromInt(3)
	//rate := decimal.RequireFromString("0.08875")
	//subtotal := price.Mul(quantity)
	//tax := subtotal.Mul(rate)
	//total := subtotal.Add(tax).Round(2)
	//fmt.Println(subtotal) // 408.06
	//fmt.Println(tax)
	//fmt.Println(total)
	//
	//var a float32 = -1.1
	//fmt.Println(a)

	//var num3 float32 = -123.0000901
	//var num4 float64 = -123.0000901
	//fmt.Println(num3)
	//fmt.Println(num4)

	//var num4 = 1.1
	//fmt.Printf("%T", num4) //float64

	//num6 := 5.12
	//num7 := .123
	//fmt.Println(num6)
	//fmt.Println(num7)

	//num8 := 5.1234e2
	//fmt.Println(num8) //512.34 科学计数法

	num9 := 5.1234e-2
	fmt.Println(num9) //0.051234 科学计数法

}
