package main

import "fmt"

func testPtr(num *int) {
	*num = 20
}

func main() {
	a := 10
	fmt.Println("调用前 a =", a) // 输出: 10

	testPtr(&a) // 把 a 的地址传进去

	fmt.Println("调用后 a =", a) // 输出: 20
}

//声明包，引包
//返回两个值
//天然支持并发
//goroutine channel支持线程间相互通信 灵魂管道
//go语句不带分号
//for循环无小括号
//行内花括号
