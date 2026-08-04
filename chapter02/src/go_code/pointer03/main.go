package main

import "fmt"

func main() {
	var num = 888
	var ptr = &num
	fmt.Printf("ptr: %d\n", ptr)
	fmt.Printf("num: %d\n", &num)

	*ptr = 333
	fmt.Printf("num: %d\n", num)

	//	1)值类型，都有对应的指针类型，
	//	形式为 *数据类型，
	//	比如，int的对应的指针就是*int，
	//	float32对应的指针类型就是*float，以此类推

	//	2）值类型包括：基本数据类型 int系列，float系列，bool、string、数组和结构体struct
	//	数组 和 结构体struct 也是值类型

}
