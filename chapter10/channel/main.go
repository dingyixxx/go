package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

var intChan chan int
var mapChan chan map[int]string //用来存放map[int]string数据
var perChan chan Person
var perChan2 chan *Person

func main() {
	//	1.channel本质是队列
	//	2.先进先出
	//	3.线程安全，多routine访问时，不需要加锁，就是说channel本身就是线程安全的
	//	4.channel有类型的，一个string的channel只能存放string类型数据

	//使用make进行初始化

	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan的值=%v\n", intChan)
	//intChan的值=0x1400009c000
	//map/channel都是引用类型
	fmt.Printf("intChan的本身的地址=%p\n", &intChan)
	//intChan的本身的地址=0x14000064038

	//为什么把channel传给另外一个函数，在另外一个函数里面操作，是同一个管道呢？
	//因为，channel是引用类型

	//向管道写入数据
	intChan <- 10 //[10]
	num := 999
	intChan <- num //[10,999]

	//看看管道的长度和cap（容量）
	//map可以自动增长，管道不可以自动增长
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))
	//channel len=2 cap=3

	<-intChan
	intChan <- 101

	//当我们给管道写入数据时，不能超过其容量
	intChan <- 50 //[10,999,50]
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))
	//channel len=3 cap=3

	//intChan <- 666 因为管道容量是3，已经有了3个了，再放一个，会怎么样
	//intChan <- 666
	//fatal error: all goroutines are asleep - deadlock!

	//取出来一个的话，看看能不能放进去新的，队头的10先出去
	var num2 int
	num2 = <-intChan ////[10,999,50] -> [999,50]
	fmt.Printf("num2-%v\n", num2)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	//在没有使用协程的情况下，如果管道数据已经全部取出，再取就会报告deaklock
	<-intChan //取出去之后，[50]
	<-intChan //取出去之后，[]
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	// 启动一个 writer 协程，1秒后往管道写入数据
	go func() {
		fmt.Println("writer: 准备写入数据...")
		time.Sleep(1 * time.Second)
		intChan <- 999
		fmt.Println("writer: 写入完成！")
	}()
	<-intChan //如果没有writer 协程（即，把上面的go func() {}()注释掉），就会立即报错

	//为什么要强调“在没有使用协程的情况下？”
	//如果有一个协程刚好在写数据，那么，阻塞等待的main协程，就不会立刻报deadlock
}
