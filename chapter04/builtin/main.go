package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1的类型%T，num1的值=%v，num1的地址%v\n", num1, num1, &num1)
	//num1的类型int，num1的值=100，num1的地址0x140000100b8

	num2 := new(int)
	//num2的类型%T => *int 是个指针
	//num2的值 => 地址1
	//num2的地址%v => 地址2
	//new给值类型分配内存，返回的是指针
	fmt.Printf("num2的类型%T，num2的值=%v，num2的地址%v，num2的指向的值%v\n", num2, num2, &num2, *num2)
	//num2的类型*int，num2的值=0x140000100d0，num2的地址0x14000064040，num2的指向的值0

	*num2 += 199
	fmt.Printf("num2的指向的值%v\n", *num2) //num2的指向的值199

	//	make也是分配内存

}
