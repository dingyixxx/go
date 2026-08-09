package main

import (
	"fmt"
	"time"
)

//goroutine中使用recover，解决协程中出现panic、导致程序崩溃问题

// 函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

// 函数
func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err...", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}

func main() {

	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main")
		time.Sleep(time.Second)
	}
}
