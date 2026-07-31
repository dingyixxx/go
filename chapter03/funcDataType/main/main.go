package main

import "fmt"

func getSum(num1 int, num2 int) int {
	return num2 + num1
}

//func myFun(fun func(int, int) int, num1 int, num2 int) int {
//	return fun(num1, num2) * 2
//}

func myFun(fun MySumFuncType1, num1 int, num2 int) int {
	return fun(num1, num2) * 2
}
func main() {
	//	函数，也是一个数据类型
	a := getSum
	fmt.Printf("a的类型是%T,getSum的类型是%T\n", a, getSum)
	//a的类型是func(int, int) int,getSum的类型是func(int, int) int

	res := getSum(10, 33)
	fmt.Println(res)

	res1 := myFun(getSum, 10, 20)
	fmt.Println(res1)
}
