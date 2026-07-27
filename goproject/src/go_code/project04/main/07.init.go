package main

import (
	"fmt"
	"unsafe"
) //比单行单行地引入，简单一些

func main() {
	var a uint = 8900
	fmt.Print(a)

	//	rune=int32
	//byte是uint8，经常放字母

	//	int默认是int64
	//	有符号 和 无符号 的 区别是：如果你想表达一个整数，希望它不带符号，希望它表示的范围大一点，则选择无符号的

	//golang的整型默认是int
	//cannot use 9223372036854775808 (untyped int constant) as int value in assignment (overflows)
	var b = 1
	b = 9223372036854775807
	fmt.Println(b)

	//	怎么查看一个变量的数据类型 fmt.Printf 变量的格式化输出
	fmt.Printf("%T \n", b)       //int
	fmt.Printf("数据类型是：%T \n", a) //uint

	var n2 int64 = 10                              //%d 十进制整数
	fmt.Printf("占用的字节数是：%d \n", unsafe.Sizeof(n2)) //占用的字节数是：8

	//	浮点型（小数类型）

}
