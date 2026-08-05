package main

import "fmt"

type T interface {
}

func main() {
	//	空接口interface{}没有任何方法, 所以, 所有的类型都实现了空接口, 它是一种单独的数据类型
	//	什么数据类型, 都可以赋值给空接口
	//	是个筐, 什么都装

	var a interface{}
	a = 1
	fmt.Println(a)

	var t T
	t = "hello"
	fmt.Println(t)

	//	interface是引用类型.如果没有赋值,默认为nil.

}
