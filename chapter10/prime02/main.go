package main

import (
	"fmt"
	"strconv"
	"time"
)

// 需求：要求统计1-20000 的数字中，哪些是素数？
// 这个问题在本章开篇就提出了，现在我们有goroutine和channel的知识后，就可以完成了 [测试数据: 20000]
// 分析思路：
// ➤ 传统的方法，就是使用一个循环，循环的判断各个数是不是素数【ok】。
// ➤ 使用并发/并行的方式，将统计素数的任务分配给多个(4个)goroutine去完成，完成任务时间短。
// 1. 画出分析思路 2. 代码实现
// 说明：使用goroutine完成后，可以在使用传统的方法来统计一下，看看完成这个任务，各自耗费的时间是多少?[用map保存primeNum]
var count int = 20000
var routineCount int = 4
var intChan chan int = make(chan int, count)
var primeChan chan int = make(chan int, count/2+1)
var doneRtChan chan int = make(chan int, routineCount) //4个协程， 分头去处理
var putStart time.Time
var checkStart time.Time

func main() {
	totalStart := time.Now()
	go putNum()

	for i := 0; i < routineCount; i++ {
		checkStart = time.Now()
		go checkIsPrimeAndPut(i)
	}
	for i := 0; i < routineCount; i++ {
		<-doneRtChan
	}
	checkEnd := time.Now()
	close(primeChan)
	fmt.Print("素数是：")
	for v := range primeChan {
		fmt.Print(strconv.Itoa(v) + ",")
	}
	totalEnd := time.Now()

	fmt.Printf("\n\n===== 耗时统计 =====\n")
	fmt.Printf("putNum 耗时:       %v\n", checkStart.Sub(putStart))
	fmt.Printf("checkIsPrime 耗时: %v\n", checkEnd.Sub(checkStart))
	fmt.Printf("总耗时:            %v\n", totalEnd.Sub(totalStart))
}

func checkIsPrimeAndPut(i int) {
	fmt.Printf("协程-%v正在判断是否为素数，并将素数放入队列\n", i)
	for {
		num, ok := <-intChan
		if ok == true {
			var isPrime bool = checkIsPrime(num)
			if isPrime == true {
				fmt.Printf("协程-%v发现素数-%v\n", i, num)
				primeChan <- num
			}
		} else {
			fmt.Printf("协程-%v都处理完了，结束\n", i)
			doneRtChan <- i
			break
			//有break，必有close，
			//记死了，
			//不可以“不close就直接for里盼false再break”，
			//因为根本监听不到，走不到这里
		}
	}

}

func checkIsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func putNum() {
	putStart = time.Now()
	for i := 0; i < count; i++ {
		intChan <- i + 1
		fmt.Printf("把num-%v给放到channel里面\n", i+1)
	}
	close(intChan)
}
