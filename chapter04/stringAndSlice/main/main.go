package main

import "fmt"

func main() {
	//	string底层是一个byte数组，java的string底层是一个char数组
	//string可以进行切片处理

	email := "laojuju@laochanchan.com"
	email += "aaaaaaaaaaa"
	str2 := email[8:]
	fmt.Println("email是：", email)
	fmt.Printf("slice1的类型是：%T\n", str2) //string，只读的类切片结构，不是数组也不是切片，是独立类型。

	//email[0]=67 string 是不可变的，直接改变数组下标对应的值，是不可以的，Cannot assign to email[0]
	//str2[0]++//Go 的 string 是不可变的，所以str2[0]也不能改
	fmt.Println("str2是：", str2)

	slice2 := []byte(email[8:])           //先转成切片
	fmt.Printf("slice2的类型是：%T\n", slice2) //[]uint8
	slice2[0] += 4
	slice2[1] += 4
	slice2[2] += 4

	fmt.Println(string(slice2)) //string入参byte array
	for _, b := range slice2 {
		fmt.Printf("%c", b)
	}
	fmt.Println()
	//peschanchan.comaaaaaaaaaaa

	email2 := "laojuju@laochanchan.com"
	//slice2[2]='婵'
	//''婵'' (type rune) cannot be represented by the type byte
	//byte这个数据类型，放不下中文
	slice3 := []rune(email2[8:]) //先转成切片
	slice3[2] = '婵'
	slice3[3] = '宝'
	slice3[4] = '虫'
	slice3[5] = '任'
	slice3[6] = '务'
	slice3[7] = '编'
	slice3[8] = '排'
	slice3[9] = '妙'

	//la婵宝虫任务编排妙n.com
	fmt.Println(string(slice3))

}
