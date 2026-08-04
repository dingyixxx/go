package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 int32 = 12 + 256
	var n3 int8
	var n4 int8
	var n5 = 128
	n3 = int8(int(n1) + n5) //-116 越界 溢出。 如果不引入变量n5，128编译的时候，都过不去，128编译的时候直接检查术来overflow。
	n4 = int8(n1) + 127     // -117 越界 溢出

	fmt.Println(n3)
	fmt.Println(n4)

	var num3 = 88
	fmt.Printf("%d", unsafe.Sizeof(num3)) //8

}
