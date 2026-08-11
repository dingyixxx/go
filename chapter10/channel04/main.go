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

	//var intChan = make(chan int, 2)
	//如果管道如果不够长(例如,只能放两个),
	//并不会报错,
	//只是会阻塞写入,
	//即, 只要调度器发现有协程还在写, 这些任务都会被挂起

	//var intChan = make(chan int, 10)
	var exitChan = make(chan bool, 1)

	//读和写，一定要用上“协程”做的
	//由“调度器”自动协调多个“用户态线程”以提高任务执行的并发度
	go readData(intChan, exitChan)
	time.Sleep(time.Second * 3)
	go writeData(intChan)
	//先读，再写，为什么也可以
	//因为readData里，只要intChan没关，读就死等

	//如果把go readData(intChan, exitChan)这行注释掉,
	//只写不读,就会因为v, ok := <-exitChan在死等
	//且"不可能有任何逻辑写入它"而报错
	//fatal error: all goroutines are asleep - deadlock!

	//那么, 有一个问题是, 为什么这次"死等"就报错了, "先读再写"也是在死等就没有报错?
	//go1.25.12/src/runtime/proc.go:443 等到进程结束了"都没等来", 那就是错误
	//8秒后, 调度器发现所有协程都"沉睡"了(没可能有东西写入管道了), 才会出现这个错误
	//fatal error: all goroutines are asleep - deadlock!
	//等到主协程结束了"等到了", 那就没有问题

	//其实,"只读不写"是同理的,如果写协程不写intChan,
	//又不关intChan,
	//读协程也是会在"go readData(intChan, exitChan)--->v, ok := <-intChan这个报错堆栈"崩溃

	//协程是主线程的守护线程，此处如果啥也不写，协程也没了
	//time.Sleep(time.Second * 15)

	for {
		mainThreadStatus() //因为exitChan只有一个元素，所以，只打印一次
		//time.Sleep(time.Millisecond * 8000)

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
			//time.Sleep(time.Millisecond * 60)
			time.Sleep(time.Millisecond * 2000)
			//把intChan的容量改10, 读速减慢, 看看会怎样?
			//会阻塞读
			//Java的LinkedBlockingQueue两把锁, 吞吐高
			//ai和我共同操作一个资源, 我读的速度慢, ai写得快

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
	for i := 0; i < 50; i++ {
		//time.Sleep(time.Millisecond * 2000)
		time.Sleep(time.Millisecond * 47)
		intChan <- i + 1
		fmt.Println("writeData写了.....下标....", i)

	}
	close(intChan)
	//写完了，就把channel一关
}

//有点像发布订阅

//channel：
//请鳖入瓮，随后，可以开始瓮中捉鳖，
//不捉则不用关（for...i...就不需要close），
//要捉（while监听到“后面没有鳖了”，就false break）则必须关（close）
