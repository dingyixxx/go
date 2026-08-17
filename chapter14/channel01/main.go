package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{}) //无缓冲区的channel，可以写在select里、接收多次，含义为“收到stopCh的信号”
	for i := 0; i < 100; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "sent result")
			case <-done:
				fmt.Println(idx, "exiting")
			}
		}(i)
	}
	//get first result
	fmt.Println("result:", <-ch)
	close(done)
	//do other work
	time.Sleep(3 * time.Second)
}
