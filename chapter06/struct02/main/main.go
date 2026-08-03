package main

import "fmt"

type Cat struct {
	Name  string
	Age   int //结构体，定义的时候，没有逗号
	Color string
	Hobby string
}

func main() {
	cat := Cat{"小花", 2, "狸花", "喝牛奶"}
	cat2 := cat
	cat2.Name = "我是cat2，我不是小花"
	fmt.Printf("cat是:%v\n", cat)
	fmt.Printf("cat.Name的地址是:%p\n", &(cat.Name))
	fmt.Printf("cat2.Name的地址是:%p\n", &(cat2.Name))
	//cat是:{小花 2 狸花 喝牛奶}
	//cat.Name的地址是:0x14000108040
	//cat2.Name的地址是:0x14000108080

	slice1 := make([]int, 0)
	slice1 = append(slice1, 1, 2, 3)
	slice2 := slice1
	slice2[0] = 999
	fmt.Printf("slice1的内存地址是:%v\n", slice1)
	fmt.Printf("slice1[0]的内存地址是:%p\n", &(slice1[0]))
	fmt.Printf("slice2[0]的内存地址是:%p\n", &(slice2[0]))
	//slice1的内存地址是:[999 2 3]
	//slice1[0]的内存地址是:0x14000016090
	//slice2[0]的内存地址是:0x14000016090

}
