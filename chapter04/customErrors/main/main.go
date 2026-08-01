package main

import (
	"errors"
	"fmt"
)

func readConfig(fileName string) (err error) {
	if fileName == "config.ini" {
		return nil
	}
	return errors.New("读取文件错误")
}

func test02() {
	err := readConfig("1config.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02...程序继续执行...")
}
func main() {
	//	go程序中，也支持自定义错误，使用errors.New和panic内置函数
	//	步骤一:errors.New("错误说明")，会返回一个error类型的值，表示一个错误
	//	步骤二:panic内置函数，接收一个interface(){}类型的值（也就是任何值了）作为参数，可以接收error类型的变量，输出错误信息，并退出程序

	//defer捕获错误
	//defer func() {
	//	//recover()
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	test02()
	fmt.Println("main...程序继续执行...")

	//	循环打印输入月份的天数

}
