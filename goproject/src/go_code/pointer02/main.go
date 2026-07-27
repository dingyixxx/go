package main

import "fmt"

func main() {
	//	1）对于基本数据类型，变量存的就是值，也叫值类型
	//	基本数据类型在内存布局
	var i int = 10 //println+tab
	//i的地址是什么
	fmt.Println("i:", &i) //值类型在内存的布局
	//	下面的var ptr *int = &i
	//1.ptr是一个指针变量
	//2.ptr的类型*int
	//3.ptr本身的值&i

	//2）获取变量的地址，用&，比如：var num int，获取num的地址：&num
	//3）指针类型，指针变量存的是个地址，这个地址指向的空间存的才是值，比如：var ptr *int=&num
	//4）获取指针类型所指向的值，使用：*
	//比如，var ptr *int ，使用*ptr获取p指向的值

	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr的指向的值=%v", *ptr)


}
