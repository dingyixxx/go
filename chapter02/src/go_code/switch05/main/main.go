package main

import "fmt"

func main() {
	//var age = 10
	//switch age {
	//case 10:
	//	fmt.Println("10...")
	//case 20:
	//	fmt.Println("20...")
	//default:
	//	fmt.Print("default...")
	//}

	//var age = 10
	////switch后面不写变量，case后面跟着表达式
	//switch {
	//case age == 10:
	//	fmt.Println("10...")
	//case age == 20:
	//	fmt.Println("20...")
	//default:
	//	fmt.Print("default...")
	//}

	var age = 10
	//case后面还可以进行范围判断
	switch {
	case age > 20:
		fmt.Println(">20...")
	case age > 5:
		fmt.Println(">5...")
	default:
		fmt.Print("default...")
	}

	//switch后也可以直接声明/定义一个变量，分号结束，不推荐
	switch grade := 3; {
	case grade > 2:
		fmt.Println("grade > 2...")
	case grade > 1:
		fmt.Println("grade > 1...")
	default:
		fmt.Print("default...")
	}

}
