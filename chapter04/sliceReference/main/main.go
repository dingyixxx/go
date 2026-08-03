package main

import "fmt"

func main() {
	var slice []int
	var arr = [...]int{1, 2, 3, 4, 5}
	slice = arr[:]     // slice 指向 arr 的底层数组
	var slice2 = slice // slice2 和 slice 指向同一个底层数组
	slice2[0] = 10     // 修改 slice2，会影响 slice 和 arr

	//slice2 [10 2 3 4 5]
	//slice [10 2 3 4 5]
	//arr [10 2 3 4 5]
	//切片是引用传递的
	fmt.Println("slice2", slice2)
	fmt.Println("slice", slice)
	fmt.Println("arr", arr)

	//数组是值传递的
	//但数组和切片，会相互改变，神奇
}
