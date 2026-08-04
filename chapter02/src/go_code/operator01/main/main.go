package main

import "fmt"

func main() {

	//var num1 byte = 'A'
	//var num2 byte = 'Z'
	//var num3 byte = 'a'
	//var num4 byte = 'z'
	//fmt.Println(num1) //65
	//fmt.Println(num2) //90
	//fmt.Println(num3) //97
	//fmt.Println(num4) //122

	var num = 200
	var ptr = &num
	fmt.Println(&num)
	fmt.Println(*ptr)

	//fmt.Scanln() //姓名，年龄，薪水，是否通过考试
	//fmt.Scanf()

}
