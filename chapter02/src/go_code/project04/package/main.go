package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//	bool类型只允许取值true和false，不能取1或者-1
	var a bool = true //等于1是不可以
	fmt.Println(a)

	var b bool = true
	fmt.Printf("占用的字节数是：%d", unsafe.Sizeof(b)) //%d 1

}
