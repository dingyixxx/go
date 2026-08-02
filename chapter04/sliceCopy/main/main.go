package main

import "fmt"

func main() {
	//切片的拷贝
	var src = []int{1, 2, 3, 4, 5}

	var slice = make([]int, 10)
	fmt.Println(slice)

	//[0 0 0 0 0 0 0 0 0 0]
	//	[1 2 3 4 5 0 0 0 0 0]
	copy(slice, src) //必须都是切片类型，不能是数组，即，切片和数组不能相互拷贝
	fmt.Println("slice=", slice)

	src[1] = 12345
	//slice= [1 2 3 4 5 0 0 0 0 0]
	//src= [1 12345 3 4 5] 拷贝完了就完了，改变切片，不会影响其“拷贝来自于的”或者“被拷贝成的”切片
	fmt.Println("slice=", slice)
	fmt.Println("src=", src)

	fmt.Println()
	var source []int = []int{1, 2, 3, 4, 5}
	var dest = make([]int, 1)
	fmt.Println("source=", source)
	copy(dest, source)
	fmt.Println("dest=", dest)
	//source= [1 2 3 4 5]
	//dest= [1]
	//容量只有1，拷贝不到

}
