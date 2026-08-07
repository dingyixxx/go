package main

import (
	"fmt"
	"time"
)

var count int = 2000
var numChan chan int = make(chan int, count)
var resChan chan map[int]int = make(chan map[int]int, count)
var doneChan chan int = make(chan int, count) //信号量

func main() {
	for i := 0; i < 8; i++ {
		go getNum(i)
	}
	time.Sleep(time.Second * 3)
	go putNum()

	for i := 0; i < 2000; i++ {
		<-doneChan
	}
	close(resChan)

	//用 for { <-ch } + ok 判断退出时，必须 close channel。用 for i := 0; i < N; i++ { <-ch } 精确计数时，不需要 close。

	fmt.Println("resChan可以开始遍历")
	for m := range resChan {
		for key, value := range m {
			fmt.Printf("res[%v]=%v\n", key, value)
		}
	}

}

func getNum(no int) {
	for {
		num, ok := <-numChan
		if ok == true {
			time.Sleep(time.Millisecond * 20)
			fmt.Printf("我是协程%v号，我现在从numChan取出一个值，%v\n", no, num)
			calcSumForN(num)
			doneChan <- 1
		} else {
			fmt.Printf("我是协程%v号，numChan取不到了\n", no)
			break
		}
	}
}

func calcSumForN(num int) {
	sum := 0
	for i := 0; i <= num; i++ {
		sum += i
	}
	time.Sleep(time.Millisecond * 10)
	var m map[int]int = make(map[int]int, 1)
	m[num] = sum
	resChan <- m
}

func putNum() {
	//fori tab快捷键，和idea一样的
	for i := 1; i <= count; i++ {
		time.Sleep(time.Millisecond * 15)
		numChan <- i * 2
	}
	close(numChan)
}
