package main

import "fmt"

func main() {
	//var age byte
	//fmt.Println("请输入你的年龄")
	//fmt.Scanln(&age) //输入年龄
	//if age > 18 {
	//	fmt.Println("你是成年人,要为自己的行为负责...")
	//} else {
	//	fmt.Println("青少年")
	//}

	if count := 30; count > 20 {
		fmt.Println("hello")
	} else if count < 10 {

	} else {
		fmt.Println("...count<20...")
	}
}
