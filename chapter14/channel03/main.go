package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
			fmt.Println("收到了:", <-ch) //写会被阻塞
		}(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("3秒过去了，读取方还在阻塞...")
	//get first result
	fmt.Println("发出去了result:", <-ch) //读会被阻塞    但因为"永远也不能写"，所以证明不了，不纠结了
	//do other work
	time.Sleep(10 * time.Second)
}
