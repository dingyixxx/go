package main

import "fmt"

func main() {
	// 类型断言的其它案例
	var x interface{}
	var b float32 = 1.1
	x = b // 空接口，可以接收任意类型

	// x => float32 [使用类型断言]
	//y := x.(float32)
	// if ok == true 可以简写为 if ok
	if y, ok := x.(float64); ok {
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	} else {
		fmt.Println("convert fail")
	}

	//fmt.Printf("y 的类型是 %T 值是=%v", y, y)
}
