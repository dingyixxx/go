package main

import "fmt"

func main() {
	//total := 97
	//fmt.Printf("%d星期%d天\n", total/7, total%7)
	//
	//var htemp float64 = 134.2// 134.2 在二进制中无法精确表示（0.2 是无限循环小数），误差会在计算中累积
	//var stemp float64 = ((htemp - 100) * 5) / 9
	//fmt.Println(stemp)

	//total := 97
	//fmt.Printf("%d星期%d天\n", total/7, total%7)
	//
	var htemp float64 = 134.2 // 134.2 在二进制中无法精确表示（0.2 是无限循环小数），误差会在计算中累积
	var stemp float64 = ((htemp - 100) * 5) / 9
	fmt.Println(stemp)

	//total := 97
	//fmt.Printf("%d星期%d天\n", total/7, total%7)

	//htemp := decimal.NewFromFloat(134.2)
	//stemp := htemp.Sub(decimal.NewFromInt(100)).Mul(decimal.NewFromInt(5)).Div(decimal.NewFromInt(9))
	//fmt.Println(stemp)
}
