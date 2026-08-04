package main

import "fmt"

func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

// 定义全局变量
var (
	dingyi = "dingyi"
	you    = "youDearYouMyFirstAndForeverAndSoReal"
)

func main() {
	sum, sub := getVal(3, 1)
	fmt.Println(sum)
	//sub="hello"
	fmt.Println(sub)

	sum2, _ := getVal(3, 44)
	fmt.Println(sum2)

	a := 0
	fmt.Println(a)

	var b int
	//b = 1
	fmt.Println(b)

	fmt.Println(dingyi)
	fmt.Println(you)

	c, d := 99, 10
	fmt.Println(c)
	fmt.Println(d)

	//多变量声明
	var tiny1 float32
	tiny1 = 1.1
	fmt.Println(tiny1)
}
