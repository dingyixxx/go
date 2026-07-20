package main

import "fmt"

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	// 方式1：用多个变量接收
	sum, sub := getSumAndSub(10, 3)
	fmt.Println("和 =", sum, "差 =", sub) // 输出: 和 = 13 差 = 7

	// 方式2：只接收其中一个值，另一个用 _ 忽略
	sum2, _ := getSumAndSub(10, 3)
	fmt.Println("和 =", sum2) // 输出: 和 = 13
}

//切片 动态数组

//defer 延迟执行

//目录结构的说明

//go的目录结构：goproject-src-go_code
//--project01
//-----main
//-----package
//--project02
//-----main
//-----package
