package main

import "fmt"

func main() {
	var address = "北京长城 110 hello world"
	fmt.Println(address)

	//	1.golang统一使用utf-8编码，中文乱码问题就不会影响程序员了
	//2.字符串一旦赋值了，就不能修改了，在go中字符串是不可变的
	var str string = "abc"
	//str[0]='b'//Cannot assign to str[0]
	fmt.Println(str)
	//	go语言的字符串的表达形式有两种：1.双引号(遇到特殊字符得转译，例如\n) 2.反引号
	fmt.Println(`aa
              aaaaaa`)

	fmt.Println(`package main

import "fmt"

func main() {
	var address = "北京长城 110 hello world"
	fmt.Println(address)

	//	1.golang统一使用utf-8编码，中文乱码问题就不会影响程序员了
	//2.字符串一旦赋值了，就不能修改了，在go中字符串是不可变的
	var str string = "abc"
	//str[0]='b'//Cannot assign to str[0]
	fmt.Println(str)
	//	go语言的字符串的表达形式有两种：1.双引号(遇到特殊字符得转译，例如\n) 2.反引号
	fmt.Println()

}
`)

	var str1 = "hello" +
		"world"
	fmt.Println(str1) //当一个拼接的操作很长时，怎么办？
	// 可以换行，加号+留在上一行

}
