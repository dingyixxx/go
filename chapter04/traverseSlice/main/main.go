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

	//切片可以动态增长 append扩容

	//append可以追加多个变量
	slice5 := []int{100, 200, 300}
	slice5 = append(slice5, 400, 500) //快捷键slice5.aapend tab
	fmt.Println(slice5)

	slice6 := append(slice5, 8888, 9999)
	//slice5= [100 200 300 400 500]
	//slice6= [100 200 300 400 500 8888 9999]
	//因此，append是不改变原切片的
	fmt.Println("slice5=", slice5)
	fmt.Println("slice6=", slice6)

	//append可以追加切片，用 切片名字... 展开
	slice7 := []int{111, 333}
	slice8 := []int{222, 444}
	slice9 := append(slice7, slice8...)
	//slice7= [111 333]
	//slice8= [222 444]
	//slice9= [111 333 222 444]
	fmt.Println("slice7=", slice7)
	fmt.Println("slice8=", slice8)
	fmt.Println("slice9=", slice9)

	//切片append操作的底层原理分析：
	//1.切片append操作的本质，是对数组扩容
	//2.go底层会创建一个新的数组newArr（安装扩容后大小）
	//3.将slice原来包含的元素拷贝到新的数组newArr
	//4.slice重新引用到newArr
	//5.注意，newArr是在底层来维护的，程序员不可见

	//	但我的“感觉”，应该是原来的capicity不够，才会创建新的数组？

}
