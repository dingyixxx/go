package main

import "fmt"

/*
转义字符示例
*/
func main() {
	fmt.Println("天龙八部雪山飞狐\r张飞厉害")
	fmt.Println("aaaaabbbbbbb北京北京北京dasdsadasdasdas路\r径: C:\\Users") // 输出: 路径: C:\Users
	fmt.Println("Hello\nWor\\ld")                                   // 换行输出
	fmt.Println("他说\"你好\"")                                         // 输出: 他说"你好"
	fmt.Println("\t\t他说tab")                                        // 输出: 他说"你好"
	fmt.Println("aaa")                                              // aaa

}
