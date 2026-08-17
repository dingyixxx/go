package main

import (
	"fmt"
)

// ❌ 错误写法：defer 在函数末尾执行，资源堆积
func processBad() {
	for i := 0; i < 3; i++ {
		fmt.Printf("循环%d: 打开资源\n", i)
		defer fmt.Printf("循环%d: 关闭资源\n", i) // 不会在这里执行！
	}
	fmt.Println("函数结束")
	// 所有 defer 在这里才执行
}

// ✅ 正确写法：用匿名函数包裹
func processGood() {
	for i := 0; i < 3; i++ {
		func() {
			fmt.Printf("循环%d: 打开资源\n", i)
			defer fmt.Printf("循环%d: 关闭资源\n", i) // 匿名函数结束时执行
			//	解决这个问题的一个方法是把代码块写成一个函数。
		}()
	}
	fmt.Println("函数结束")
}

func main() {
	fmt.Println("=== 错误写法 ===")
	processBad()

	fmt.Println("\n=== 正确写法 ===")
	processGood()
}
