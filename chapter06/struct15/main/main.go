package main

import (
	"fmt"
)

// 如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出
type integer11 int

func (self integer11) String() string {
	return "我是修改之后的..." + string(self)
}

func main() {
	var num integer11 = 999
	fmt.Println(num)

}
