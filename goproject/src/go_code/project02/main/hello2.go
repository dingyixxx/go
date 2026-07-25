package main

import "fmt"

// format包,go定义的变量或者import的包如果没有被用到，则不能编译通过
// 行内花括号, go设计者: 一个问题尽量只有一个解决方法
// go的入口函数是main "fmt" imported and not used
func main() { //func表示后面是一个函数
	var b = 99
	fmt.Println(b)
	a := 10 // 用 _ 接收，表示"我不用这个值"
	//fmt.Println(a)
	fmt.Println(a)
	//declared and not used: a

	//+加号拼接,和,逗号拼接,输出时都不会真的换行
	fmt.Println("Hello, World!Hello, World!Hello, World!1111" +
		"222222Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!H333333" +
		"4444ello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hell555555" +
		"6666o, World!Hello, World!Hello, World!Hello, World!Hello, World!Hel" +
		"lo, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!")

	fmt.Println("aaaaaa", "bbbbbb", "cccccc", "dddddd")

	b = 3 + 4
	fmt.Println(b)

	_ = 999 // 用 _ 接收，表示"我不用这个值"
	//	ctrl+alt+l 和idea一样的
	//gofmt
	//	加号+前后各加一个空格

}
