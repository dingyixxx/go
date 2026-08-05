package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{} // 空接口，可以存任何类型
	var point Point = Point{1, 2}
	a = point // ✅ ok，Point 赋值给空接口

	// 如何将 a 赋给一个 Point 变量？
	var b Point
	//b = a // ❌ 编译报错！不能直接赋值
	b = a.(Point) // 可以了，类型断言

	fmt.Println(b)
}
