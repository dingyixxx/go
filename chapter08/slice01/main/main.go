package main

import "fmt"

func main() {
	//如何从切片中删除元素
	var slice1 []int = []int{1, 4, 5, 66, 89, 21, 43}
	slice1 = append(slice1[:3], slice1[4:]...)
	fmt.Println(slice1)
}
