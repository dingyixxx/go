package main

import (
	"fmt"
	"strconv"
	_ "unsafe" //如果没有使用到一个包，直接加下划线_表示忽略
)

func main() {
	var n1 int32 = 444
	fmt.Printf("%b\n", 444) //1 10111100
	var f = float32(n1)
	var num2 int8 = int8(n1) //444精度丢失
	fmt.Printf("%d,%f,%d\n", n1, f, num2)
	res, _ := strconv.ParseInt("10111100", 2, 64)
	//188 > 127（超出 int8 正数范围），所以：
	//[188 - 256 = -68]
	fmt.Printf("10111100-%d \n", res)

	var num3 int64 = int64(n1) //Cannot use 'n1' (type int32) as the type int64
	fmt.Println(num3)

	fmt.Printf("%t", n1) //%!t(int32=444) 被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化

	//	在转换中，比如，将int64转成int8，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样
	var num4 int64 = 999999
	fmt.Printf("%b\n", num4) //1111 01000010 00111111
	res2, _ := strconv.ParseInt("00111111", 2, 8)
	fmt.Println(res2) //63

	var num5 int8 = int8(num4)
	fmt.Println(num5)
}
