package main

import "fmt"

func main() {

	//基本数据类型

	//	数值型
	//java的 short byte int long 1 2 4 8
	//int8 int16 int32 int64
	//拿掉一位作为符号位, 所以是2的少一次方次幂
	//-2^7-(2^7-1)
	//-2^15-(2^15-1)
	//-2^31-(2^31-1)
	//-2^63-(2^63-1)

	//uint8 uint16 uint32 uint64 无符号的,表示的范围更大
	//rune是int32的别名
	var c rune
	fmt.Println(c)

	//float32 float64 double

	//字符型
	//byte(没有专门的字符型,使用byte来保存单个字母字符)
	//utf-8 一个汉字占3个字节 不能存汉字
	f := 'a'
	fmt.Println(f)

	//布尔型
	a := true
	fmt.Println(a)

	//	字符串
	str := "hello world"
	fmt.Println(str)

	fmt.Println("------")
	var i int8 = -128
	fmt.Println(i)
	var j int8 = 127
	fmt.Println(j)

	//uint8 uint16 uint32 uint64
	//0-255 0-2^16-1 0-2^32-1 0-2^64-1
	var k uint16 = 256
	fmt.Println(k)

	var m uint8 = 4
	fmt.Println(m)
	//派生数据类型
	//1.指针
	//2.数组
	//3.结构体struct类似于class
	//4.管道channel
	//5.函数(也是一种类型)
	//6.切片(slice)
	//7.接口interface
	//8.map

}
