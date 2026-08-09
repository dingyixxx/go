package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//select可以解决从管道取数据的阻塞问题

	//1.定义一个管道，10个数据int
	count := 10
	intChan := make(chan int, count)
	for i := 0; i < count; i++ {
		intChan <- i + 1
	}

	//2.定义一个管道，5个数据string
	count2 := 5
	stringChan := make(chan string, count2)
	for i := 0; i < count2; i++ {
		stringChan <- "str-" + strconv.Itoa(i+1)
	}

	//传统方法在遍历管道时，如果不关闭，会导致deadlock
	//for v := range stringChan {
	//	fmt.Println(v)
	//} //fatal error: all goroutines are asleep - deadlock!
	//在实际开发中，确定惯到什么时候能close，有时，并不是简单的事
	//此时，就可以用select
yuan:
	for {
		select {
		case v := <-stringChan:
			fmt.Println(v) //如果这个取不到，就去下面的取
		case v := <-intChan:
			fmt.Println(v)
		default:
			fmt.Println("都取不到，歇一会儿...")
			time.Sleep(time.Second * 5)
			break yuan
		}
	}
	fmt.Println("回来了...")

}
