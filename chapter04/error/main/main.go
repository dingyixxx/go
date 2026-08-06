package main

import (
	"fmt"
	"time"
)

func test() {
	//使用defer+recover来捕获和处理异常 ephemeral
	defer func() {
		//recover()
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() //匿名函数调用
	//如果没有加defer，那就彻底崩了
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Printf("res=%d", res)
}

func main() {
	//	panic 恐慌 致命错误 进程直接关闭
	//panic: runtime error: integer divide by zero
	//
	//goroutine 1 [running]:
	//main.test()
	//	/Desktop/go/chapter04/error/main/calc_test.go:8 +0x1c
	//main.main()
	//	/Desktop/go/chapter04/error/main/calc_test.go:14 +0x1c
	//
	//Process finished with the exit code 2
	test()

	//	在默认情况下，当发生错误后（panic），程序就会退出（崩溃）。
	//	如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。还可以在捕获到错误后，给管理员一个提示（邮件、短信...）

	//	defer panic recover
	//	简单描述，就是：go中可以抛出一个panic的异常，然后再defer中通过recover捕获这个异常，然后正常处理
	for {
		fmt.Println("我还在哦...")
		time.Sleep(time.Millisecond * 300)
	}
}
