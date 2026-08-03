package main

import "fmt"

// 结构体struct里面的各个属性，都不用加逗号
type Person struct {
	Name string
	Age  int
}

func main() {
	n := 10
	fmt.Printf("%v\n", &n) // 0x1400010e0f0  ← 地址

	s := "hello"
	fmt.Printf("%v\n", &s) // 0x1400010e0f8  ← 地址

	arr := [3]int{1, 2, 3}
	fmt.Printf("%v\n", &arr) // &[1 2 3]       ← 内容

	p := Person{"丁一", 35}
	fmt.Printf("%v\n", &p) // &{丁一 35}     ← 内容

	//不管什么类型，%p 永远打印地址
	fmt.Printf("%p\n", &n)   // 0x1400010e0f0
	fmt.Printf("%p\n", &s)   // 0x1400010e0f0
	fmt.Printf("%p\n", &arr) // 0x1400010e0f0
	fmt.Printf("%p\n", &p)   // 0x1400010e0f0

}
