package main

import "fmt"

type Animal interface {
	SayHello()
	SetNickname(name string)
}

type Dog struct {
	Name string
}

func (d *Dog) SayHello() {
	fmt.Printf("%v 说 汪汪\n", d.Name)
}
func (d *Dog) SetNickname(name string) {
	d.Name = name
}

type Cat struct {
	Name string
}

func (c *Cat) SayHello() {
	fmt.Printf("%v 说 喵 \n", c.Name)
}
func (c *Cat) SetNickname(name string) {
	c.Name = name
}

type Human struct {
}

// 接口, 是引用传递
func (human Human) KeepPet(theAnimal Animal, called string) {
	theAnimal.SetNickname(called)
	theAnimal.SayHello()
}

func main() {
	//	interface是引用类型.如果没有赋值,默认为nil.
	human := Human{}
	c := &Cat{} //不传引用的话, 会报编译错误
	human.KeepPet(c, "加菲")
	fmt.Println(c) //

	d := &Dog{}
	human.KeepPet(d, "哮天犬")
	fmt.Println(d)

}
