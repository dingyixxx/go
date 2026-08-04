package main

import "fmt"

//	函数是函数,方法是方法
//	之前讲的那个func是函数
//	现在讲的方法是结构体struct的实例, 调用某方法

type A struct {
	Num int
}

// 函数名前面的括号,写对象
// 这个地方的self是值传递,不是引用传递
func (self A) test() {
	self.Num = 999
	fmt.Println(self.Num)
}

func (self *A) test1() {
	(*self).Num = 10086
	fmt.Println((*self).Num)
	fmt.Printf("test1的self的A的地址:%p\n", self)
}

func main() {
	baby := A{
		Num: 386,
	}
	baby.test() //
	fmt.Println(baby.Num)
	baby.test1() //
	fmt.Println(baby.Num)
	fmt.Printf("main的baby的A的地址:%p\n", &baby)
}
