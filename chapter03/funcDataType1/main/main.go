package main

import "fmt"

func main() {
	//	go支持自定义数据类型，相当于给数据类型取了一个别名

	type myInt int //这时myInt就等价int来使用
	// 给int取了别名myInt，在go中myInt和int虽然都是int类型，但是go认为myInt和int是两个类型

	type MySumFuncType func(int, int) int //这是mySum就等价一个函数类型func(int,int)int

	var dingyiFunc MySumFuncType = func(i int, i2 int) int {
		return i + i2
	}
	fmt.Println(dingyiFunc(10, 20))

	//var n1 int=40
	//var n2 myInt
	//n2=n1 //Cannot use 'n1' (type int) as the type myInt
	//fmt.Println(n2)

	var n1 int = 40
	var n2 myInt
	n2 = myInt(n1)
	fmt.Println(n2)

}
