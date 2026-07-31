package main

import "fmt"

func swap(n1 *int, n2 *int) {
	temp := *n1 // 取出 n1 指向的值，比如 10
	*n1 = *n2   // 把 n2 指向的值，写入 n1 指向的内存
	*n2 = temp  // 把 temp 的值，写入 n2 指向的内存
}

func main() {
	num1 := 11
	num2 := 22
	swap(&num2, &num1)
	fmt.Printf("num1:%d\n", num1)
	fmt.Printf("num2:%d\n", num2)
	//	值类型的变量，原来也可以用“指针”来交换
	//	虽然，“要是我的话，我就直接初始化两个新的变量来接收结果了”
}
