package main

import "fmt"

// 自定义函数类型：接收两个int，返回一个int
type mySum func(int, int) int

func sum2(n1 int, n2 int) int {
	return n1 + n2
}

func sum22(n1, n2, n3 int) int {
	return n1 + n2
}

// 接收一个 mySum 类型的函数作为参数
func myFunc(funcVar mySum, num1 int, num2 int) int {
	return funcVar(num1, num2)
}

func main() {
	a := sum2
	//b := sum22
	fmt.Println(myFunc(a, 1, 2)) // 3
	//fmt.Println(myFunc(b, 1, 2))  // 3 sum22不能作为入参传入myFunc，因为类型不对

}
