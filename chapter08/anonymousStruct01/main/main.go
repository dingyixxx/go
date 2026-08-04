package main

import (
	"fmt"
	"go_code/chapter08/anonymousStruct01/model"
)

// 结构体 B，嵌入 A（匿名字段）
type B struct {
	model.A
	Name string
}

func (b *B) SayOk() {
	fmt.Printf("B SayOk-%q\n", b.Name)
}
func main() {
	//var b B
	//b.A.Name = "tom"
	////b.A.age = 19 // 跨包的话，小写字母开头的属性无法访问
	//b.A.SayOk()
	////b.A.hello() // 跨包的话，小写字母开头的方法无法访问
	//b.A.Hello()

	var b B
	b.Name = "b.Name赋值的名字Betty"
	b.A.Name = "b.A.Name赋值的名字Alice"
	b.NianLing = 99
	b.SayOk()
	b.A.SayOk()
	b.Hello()
	//B SayOk-"b.Name赋值的名字Betty"
	//A SayOk-"b.A.Name赋值的名字Alice"
	//A Hello-"b.A.Name赋值的名字Alice"
}
