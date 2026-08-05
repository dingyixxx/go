package main

import (
	"fmt"
	"sort"
)

func main() {
	var intSlice = []int{99, 3, 4, 1, 12, 6, 54, 14, 18, 9, 41, 84}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//	interface是引用传递

	//	对结构体切片排序

}
