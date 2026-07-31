package main

import "fmt"

var globalFun = func(n1 int, n2 int) int {
	return n1 * n2
} //globalFun是一个全局匿名函数

func main() {
	//	匿名函数
	//1.定义匿名函数的时候，就调用
	//2.将匿名函数赋值给一个变量，再通过变量调用
	//3.全局匿名函数

	//	int32最大值2147483647
	//	int64最大值9223372036854775807
	//fmt.Println(math.MaxInt64)
	//fmt.Println(math.MaxInt32)
	//7668574880791085844
	//9223372036854775807

	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(33, 109)
	fmt.Println(res1)

	dingyiFunc := func(n1 int, n2 int) int {

		res1 := func(n1 int, n2 int) int {
			return n1 + n2
		}(33, 109)
		fmt.Println(res1) //函数里面套着匿名函数

		return n1 + n2
	}

	fmt.Printf("%T\n", dingyiFunc) //func(int, int) int
	fmt.Println(dingyiFunc(99, 11))
	fmt.Println(globalFun(33, 144))

}
