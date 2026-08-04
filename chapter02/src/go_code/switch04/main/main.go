package main

import "fmt"

func main() {
	count := 111
	//n3 := 5
	switch count {
	//case可以有多个值
	case 5, 10, 100:
		fmt.Println("ok")
	//case n3:
	//	fmt.Println("ok")
	//case 5:
	//	fmt.Println("ok")
	default:
		fmt.Println("都没有命中...")
	}
}
