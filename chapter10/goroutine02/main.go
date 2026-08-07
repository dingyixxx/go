package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 3; i++ {
		fmt.Println("test() hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	fmt.Println("test() 自己执行完毕")
}

func main() {
	go test()
	for i := 1; i <= 5; i++ {
		fmt.Println("main() hello,golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
