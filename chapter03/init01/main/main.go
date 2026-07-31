package main

import "fmt"

func getNum() int {
	fmt.Println("getNum执行中......")
	return 199
}

func init() {
	fmt.Println("init...") //go的init函数，类似于java的静态代码段or普通代码段
	//Name = "丁一"
	//Age = 35
	//fmt.Printf("Name-%v\n", Name)
	//fmt.Printf("Age-%v\n", Age)
}

var Name string
var Age int = getNum()

func main() {
	//	init函数，该函数会在main函数执行前，被Go运行框架调用，也就是说init会在main函数之前被调用
	fmt.Println("main...age=", Age)
}

//如果一个文件，同时包括“全局变量定义”，init函数和main函数
//则执行的流程是：
//-变量定义
//-init函数
//-main函数

//getNum执行中......
//init...
//main...age= 199
