package main

import "fmt"

func main() {
	for {
		var num int64
		fmt.Println("请输入num这个数字：")
		fmt.Scanln(&num)
		fmt.Println("您输入的数字是：", num)
		if num == 0 {
			break
		}
	}
}
