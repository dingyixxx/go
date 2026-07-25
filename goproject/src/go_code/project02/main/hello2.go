package main

import "fmt"

//format包,go定义的变量或者import的包如果没有被用到，则不能编译通过

// go的入口函数是main "fmt" imported and not used
func main() { //func表示后面是一个函数
	var b = 99
	fmt.Println(b)
	a := 10 // 用 _ 接收，表示"我不用这个值"
	//fmt.Println(a)
	fmt.Println(a)
	//declared and not used: a
	fmt.Println("Hello, World!")

	_ = 999 // 用 _ 接收，表示"我不用这个值"

}
