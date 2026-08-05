package main

import "fmt"

type Animal interface {
	SetHelloSentence()
	SetNickname(name string)
}

type AnimalProps struct {
	Name          string
	HelloSentence string
}

type Dog struct {
	AnimalProps
}

func (d *Dog) SetHelloSentence() {
	d.HelloSentence = "汪汪"
}
func (d *Dog) SetNickname(name string) {
	d.Name = name
}

type Cat struct {
	AnimalProps
}

func (c *Cat) SetHelloSentence() {
	c.HelloSentence = "喵"
}
func (c *Cat) SetNickname(name string) {
	c.Name = name
}

type Human struct {
}

func (human Human) KeepPet(theAnimal Animal, called string) {
	fmt.Printf("theAnimal的地址是：%p\n", theAnimal)
	fmt.Printf("theAnimal的值是：%v\n", theAnimal)
	theAnimal.SetNickname(called)
	theAnimal.SetHelloSentence()
}

func main() {
	human := Human{}
	//SayHello和SetNickname方法定义的时候是指针，则必须传指针, 否则KeepPet会报编译错误
	c := &Cat{}
	fmt.Printf("c的地址是：%p   ", c)
	//没接口的话，不传指针，即c := Cat{}，下面两行也不报编译错误，而且会自动找到指针
	//c.SetNickname("加菲")
	//c.SetHelloSentence()
	human.KeepPet(c, "加菲")
	fmt.Printf("%q的打招呼语是%q\n", (*c).Name, (*c).HelloSentence)
	fmt.Println("------------------------")
	d := &Dog{}
	fmt.Printf("d的地址是：%p   ", d)
	human.KeepPet(d, "哮天犬")
	fmt.Printf("%q的打招呼语是%q\n", (*d).Name, (*d).HelloSentence)
}

//python的多态，“方法定义的时候，接收猫，但实际可以传狗”
//golang不可以，如果方法定义接收猫，传狗则编译报错

//---
//方法调用，能传指针传指针
