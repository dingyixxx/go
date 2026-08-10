package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	for i := 1; i <= 3; i++ {
		fmt.Println("main() hello,golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	//	协程，是主线程的守护线程
}
