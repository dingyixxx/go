package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

type integer111 int

func (num integer111) Say() {
	fmt.Println("num --- Say()")
}

// 输入func选第二个然后tab快捷键
func (s Stu) Say() {
	fmt.Println("Stu Say()")
}

func main() {
	//var a AInterface
	//a.Say() //编译虽然不报错,但运行的时候就报错
	//接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
	//要实现所有的方法

	var stu Stu
	var b AInterface = stu
	b.Say()

	var i integer111 = 10
	var c AInterface = i
	c.Say()

}
