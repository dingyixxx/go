package main

import "fmt"

func main() {
	var n1 int32 = 20
	var n2 int32 = 20
	switch n1 {
	//switch和case的数值类型，要相等，例如，不能一个是int64，一个是int32，否则编译都通不过
	case n2:
		fmt.Println("ok1...n2...第一个")
	case n2:
		fmt.Println("ok1...n2...第二个")
	case 20:
		fmt.Println("ok2...20...第一个") //case后面的常量不能重复，变量可以重复...
	//case 20:
	//	fmt.Println("ok2...20...第二个")
	default:
		fmt.Println("没有匹配到...")
	}
}
