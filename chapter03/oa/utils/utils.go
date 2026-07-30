package dabaobei //包名 和 文件夹名 也可以不一致，此处，就不一致了

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("msin")
}

var Num1 int = 100
var num2 int = 22

func Cal() {
	fmt.Println("cal")
}

// 首字母大写
func SayHello() int {
	fmt.Println("sayHello")
	return 99
}

func Calc1(num1 float64, num2 float64, operator byte) float64 {
	num1d := decimal.NewFromFloat(num1)
	num2d := decimal.NewFromFloat(num2)

	switch operator {
	case '+':
		fmt.Println("加法")
		fmt.Println(num1d.Add(num2d))
		f, _ := num1d.Add(num2d).Float64()
		return f
	case '-':
		fmt.Println("减法")
		fmt.Println(num1d.Sub(num2d))
		f, _ := num1d.Sub(num2d).Float64()
		return f
	case '*':
		fmt.Println("乘法")
		fmt.Println(num1d.Mul(num2d))
		f, _ := num1d.Mul(num2d).Float64()
		return f
	case '/':
		fmt.Println("除法")
		f, _ := num1d.DivRound(num2d, 8).Float64()
		return f
	default:
		fmt.Println("操作有误...")
	}
	return 0.0
}
