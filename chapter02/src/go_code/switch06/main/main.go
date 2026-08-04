package main

import "fmt"

func main() {
	//switch的穿透，fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("10")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("20")
		fallthrough //默认只能穿透一层
	case 30:
		fmt.Println("30")
	default:
		fmt.Println("default")
	}

}
