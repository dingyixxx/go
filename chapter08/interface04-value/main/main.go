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

func (d Dog) SetHelloSentence() {
	d.HelloSentence = "汪汪"
}
func (d Dog) SetNickname(name string) {
	d.Name = name
}

type Cat struct {
	AnimalProps
}

func (c Cat) SetHelloSentence() {
	c.HelloSentence = "喵"
}
func (c Cat) SetNickname(name string) {
	c.Name = name
}

type Human struct {
}

func (human Human) KeepPet(theAnimal Animal, called string) {
	fmt.Printf("theAnimal的地址是：%p\n", &theAnimal)
	fmt.Printf("theAnimal的值是：%v\n", theAnimal)
	theAnimal.SetNickname(called)
	theAnimal.SetHelloSentence()
}

func main() {
	human := Human{}
	c := &Cat{}
	fmt.Printf("c的地址是：%p   ", &c)
	human.KeepPet(c, "加菲")
	fmt.Printf("%q的打招呼语是%q\n", c.Name, c.HelloSentence)
	fmt.Println("------------------------")
	//传指针也不报错,但没有用
	d := &Dog{}
	fmt.Printf("d的地址是：%p   ", &d)
	human.KeepPet(d, "哮天犬")
	fmt.Printf("%q的打招呼语是%q\n", d.Name, d.HelloSentence)
}
