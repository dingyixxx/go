package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	基本数据类型和string的转换

	// 基本数据类型 转 string
	var num1 = 99
	var num2 = 23.456
	var b = true
	var myChar byte = 'h'
	var str string //空的str

	//	使用第一种方式来转换
	str1 := fmt.Sprintf("%d", num1)                  //%d十进制 %b二进制
	fmt.Printf("typeOf=%T, value=%q \n", str1, str1) //格式化输出，而不是println
	str2 := fmt.Sprintf("%f", num2)
	fmt.Printf("typeOf=%T, value=%v \n", str2, str2)
	str3 := fmt.Sprintf("%t", b)                     //%t专门转换bool
	fmt.Printf("typeOf=%T, value=%q \n", str3, str3) //%q带引号
	str4 := fmt.Sprintf("%c", myChar)
	fmt.Printf("typeOf=%T, value=%v \n", str4, str4)

	fmt.Println(str4)
	fmt.Println(str)

	num3, _ := strconv.Atoi("01101")
	fmt.Println(num3)

	num4, _ := strconv.ParseInt("01101", 2, 64)
	fmt.Println(num4)

}
