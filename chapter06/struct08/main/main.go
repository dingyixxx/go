package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {
	var a A
	var b B
	a = A(b)
	//类型相同，才可以强转
	//名字相同，才可以强转
	//属性不多不少，才可以强转

	fmt.Println(a, b)
}
