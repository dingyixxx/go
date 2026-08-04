package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	//很好，goland也有cmd+alt+t 用while来包裹该代码段
	for {
		var num1 float64
		var num2 float64
		var operator byte //+-*/ 用byte
		fmt.Println("请输入两个数字和一个运算符+-*/：")
		fmt.Scanf("%f %f %c\n", &num1, &num2, &operator)
		fmt.Printf("您输入的运算符是：%f,%f,%c\n", num1, num2, operator)

		num1d := decimal.NewFromFloat(num1)
		num2d := decimal.NewFromFloat(num2)

		switch operator {
		case '+':
			fmt.Println("加法")
			fmt.Println(num1d.Add(num2d))
		case '-':
			fmt.Println("减法")
			fmt.Println(num1d.Sub(num2d))
		case '*':
			fmt.Println("乘法")
			fmt.Println(num1d.Mul(num2d))
		case '/':
			fmt.Println("除法")
			fmt.Println(num1d.DivRound(num2d, 8))

		default:
			fmt.Println("操作有误...")

		}
	}
}
