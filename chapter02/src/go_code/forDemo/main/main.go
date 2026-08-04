package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	count := 5
	var score float64
	sum := decimal.NewFromFloat32(0.0)
	for i := 0; i < count; i++ {
		fmt.Printf("请输入第%d个学生的成绩:\n", i+1)
		fmt.Scanln(&score)
		fmt.Printf("第%d个学生的成绩是：%v\n", i+1, score)
		sum = sum.Add(decimal.NewFromFloat(score))
		//如果不写sum=sum.Add(decimal.NewFromFloat(score))，
		//而是写成sum.Add(decimal.NewFromFloat(score))，
		//那肯定是不对的
	}
	avg := sum.DivRound(decimal.NewFromInt(5), 4)

	//count := 5
	//var input string
	//sum := decimal.NewFromFloat32(0.0)
	//for i := 0; i < count; i++ {
	//	fmt.Printf("请输入第%d个学生的成绩:\n", i+1)
	//	fmt.Scanln(&input)
	//	score, _ := decimal.NewFromString(input)
	//	fmt.Printf("第%d个学生的成绩是：%v\n", i+1, score)
	//	sum = sum.Add(score)
	//}
	//avg := sum.DivRound(decimal.NewFromInt(5), 4)

	fmt.Println(sum)
	fmt.Print(avg)

}
