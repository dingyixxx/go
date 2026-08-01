package main

import "fmt"

func main() {
	slice1 := []int{11, 22, 33, 44, 55, 66}
	for i := range slice1 {
		fmt.Println(slice1[i])
	}
	fmt.Println()
	for i, i2 := range slice1 {
		fmt.Println(i, "-", i2)
	} //和idea的快捷键一样，也适用xx.for
	fmt.Println()

	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	numArr := [5]int{1, 2, 3, 4, 5}
	slice := numArr[:3]
	fmt.Println(slice)

	slice2 := numArr[3:]
	fmt.Println(slice2)

	slice3 := numArr[:]
	fmt.Println(slice3)
	//	初始化切片时，如果start是0或者end是len，那么，可以简写
	//slice :=numArr[start:end]
	//slice :=numArr[:end]
	//slice :=numArr[start:]
	//slice :=numArr[:]

	//	切片定义完后，需要引用到一个数组或者make一个空间供切片来使用

	//切片可以继续切片

	slice4 := slice[1:3]
	slice[1] = 22222222
	slice4[1] = 33333333333333333
	//[22222222 33333333333333333]
	//[1 22222222 33333333333333333 4 5]
	fmt.Println(slice4)
	fmt.Println(numArr)
}
