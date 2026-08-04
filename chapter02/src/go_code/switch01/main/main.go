package main

import "fmt"

func main() {

	//同时声明两个变量
	x1, x2 := 11, 12
	fmt.Println(x1)
	fmt.Println(x2)

	//	golang的switch后面，不需要加break
	count := 50
	switch count {
	case 100:
		fmt.Println("等于100")
	case 20:
		fmt.Println("等于20")
	case 50:
		fmt.Println("等于50")
	default:
		fmt.Println("default...")
	}
}
