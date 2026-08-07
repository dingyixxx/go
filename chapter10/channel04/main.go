package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//问题
//请完成goroutine和channel协同工作的案例，具体要求：
//1) 开启一个writeData协程，向管道intChan中写入50个整数。
//2) 开启一个readData协程，从管道intChan中读取writeData写入的数据。
//3) 注意：writeData和readData操作的是同一个管道
//4) 主线程需要等待writeData和readData协程都完成工作才能退出

// 实现思路
// 两个channel
// ┌──────────────────┐
// │    Go 主线程      │
// │                  │
// │  ┌────────────┐  │
// │  │     P      │──┼────→ ──────────────────┐
// │  ────────────┘  │      │  协程1(writeData) │──→ ┌─────────────────────┐
// │                  │      └──────────────────┘     │ channel 管道队列     │
// │  ┌────────────┐  │                               │      intChan        │
// │  │  for { }   │──┼──────────────────────────────→│  ┌───┬───┬───┬───┐  │
// │  └────────────┘  │      ┌──────────────────┐     │  │ 1 │ 2 │ 5 │ 8 │  │
// │                  │      │  协程2(readData)  │←───│  └───┴───┴───┴───┘  │
// │                  │      │                  │     └─────────────────────┘
// │                  │      │ 读完50个数据,     │
// │                  │      │ 向exitChan写入true│──→ ┌─────────────────────┐
// │                  │      │ 并关闭该管道       │     │ channel 管道队列     │
// │                  │      └──────────────────┘     │      exitChan       │
// │                  │                               │  ┌──────┐           │
// │                  │                               │  │ true │           │
// │                  │                               │  └──────┘           │
// │                  │                               └─────────────────────┘
// │                  │
// └──────────────────┘
// │
// ▼
// 程序退出

var canExit = false
var mu sync.Mutex

func main() {
	//一定要make，且一定要make指定的capacity
	var intChan = make(chan int, 50)
	var exitChan = make(chan bool, 1)

	//读和写，一定要用上“协程”做的
	//由“调度器”自动协调多个“用户态线程”以提高任务执行的并发度
	go readData(intChan, exitChan)
	time.Sleep(time.Second * 3)
	go writeData(intChan)
	//先读，再写，为什么也可以
	//因为readData里，只要intChan没关，读就死等

	//协程是主线程的守护线程，此处如果啥也不写，协程也没了
	//time.Sleep(time.Second * 15)

	for {
		mainThreadStatus() //因为exitChan只有一个元素，所以，只打印一次
		v, ok := <-exitChan
		if ok == false {
			fmt.Println("现在，可以退出了...")
			break
		}
		fmt.Println("等等看...", v)
	}

	//for {
	//	mu.Lock()
	//	if canExit == true {
	//		break
	//	}
	//	mu.Unlock()
	//	time.Sleep(time.Millisecond * 100) // 避免忙等待

	//	mainThreadStatus()
	//主线程状态: goroutine 1 [running]:

	//用全局变量，其实也是可以实现的，
	//但是要一直轮询
	//且会有Found 1 data race(s)，
	//所以，最好加上mu.Lock()

	//}

}

func mainThreadStatus() {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("主线程状态: %s\n", string(buf[:n]))
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if ok == true {
			fmt.Println("readData读到了", v)
			time.Sleep(time.Millisecond * 60)
		} else {
			fmt.Println("readData读完了...")
			break
		}
	}
	//	读完了，就告诉exitChan，然后把exitChan关了
	exitChan <- true
	close(exitChan)
	mu.Lock()
	canExit = true
	mu.Unlock()
}

func writeData(intChan chan int) {
	for i := 0; i < cap(intChan); i++ {
		time.Sleep(time.Millisecond * 47)
		intChan <- i + 1
		fmt.Println("writeData写了......", i+1)

	}
	close(intChan)
	//写完了，就把channel一关
}

//有点像发布订阅
//channel：请鳖入翁，随后，可以开始瓮中捉鳖，不捉不用关
