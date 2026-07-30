package main

import "fmt"

func main() {
	var x interface{}
	//var y = 10.0
	var y float32 = 1.33
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型～：%T", i)
	case int:
		fmt.Printf("x是int型")
	case float64:
		fmt.Printf("x是float64型")
	case func(int) float64:
		fmt.Printf("x是func(int)型")
	case bool, string:
		fmt.Printf("x是bool或string型")
	default:
		fmt.Printf("未知型")
	}

}
