package main

import "fmt"

func main() {
	//	1.channel一旦关闭，只能读，不能写
	intChan := make(chan int, 3)
	intChan <- 999
	intChan <- 888
	close(intChan)
	//intChan <- 777 //panic: send on closed channel

	for i := 0; i < cap(intChan); i++ {
		fmt.Println(<-intChan)
	}

	//	2.channel一旦关闭，就不能再开了

}
