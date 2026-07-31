package main

import "fmt"

// *星号跟着指针变量，&跟着值数据类型
func main() {
	num := 1
	test_ptr(&num)
	fmt.Printf("outer-1-%v\n", num)
	test(num)
	fmt.Printf("outer-2-%v\n", num)
}

func test(num int) {
	num++
	fmt.Printf("test-%v\n", num)
	fmt.Printf("test-内存地址-%v\n", &num) //0x14000106028
}

func test_ptr(num *int) {
	*num++
	fmt.Printf("test_ptr-%v\n", *num)
	fmt.Printf("test_ptr-内存地址-%v\n", &num) //0x14000108038
}
