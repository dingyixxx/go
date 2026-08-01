package main

import "fmt"

func main() {
	//切片有三种定义方式
	//1.引用已经创建好的数组（如同之前的案例）
	//2.make（类型，len，cap）make创建切片时，底层的数组对外不可见
	var slice = make([]int, 4)
	slice[0] = 100
	slice[1] = 200
	slice = append(slice, 100)
	fmt.Println(slice) //[100 200 0 0 100]

	//切片必须make使用，不make不行

	//	3.直接指定具体的数组
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
