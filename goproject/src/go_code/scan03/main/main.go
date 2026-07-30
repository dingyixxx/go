package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入姓名(should be 丁一)：")
	fmt.Scanln(&name)
	fmt.Println("请输入密码(should be 1234)：")
	var password string
	fmt.Scanln(&password)

	//go的字符串，是值类型，所以可以这么判断
	if name == "丁一" && password == "1234" {
		fmt.Println("你输入的用户名和密码，是正确的")
	} else {
		fmt.Println("wrong...")
	}
}
