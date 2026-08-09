package main

import "fmt"

func main() {
	//管道可以声明为只读或者只写

	//1.默认情况下，管道是双向

	//2.声明为只写
	var chan2 chan<- int //只进不出
	chan2 = make(chan int)
	chan2 <- 20
	//num :=<- chan2

	//3.声明为只读
	var chan3 <-chan int //只出不进
	//chan3 <- 20
	num := <-chan3
	fmt.Println(num)

}
