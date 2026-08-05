package main

import "fmt"

// Monkey 结构体
type Monkey struct {
	Name string
	Animal
}

type Animal struct {
	Name          string
	HelloSentence string
}

type BirdAble interface {
	Flying()
}

func SeeBlueSkyAndCloud(bird BirdAble) {
	fmt.Println("看见黄鹂")
}
func (lm *LittleMonkey) Flying() {
	fmt.Println(lm.Name, "通过研发飞行翼，会了Flying...")
}

type FishAble interface {
	Swimming()
}

func (lm *LittleMonkey) Swimming() {
	fmt.Println(lm.Name, "Swimming...")
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树..")
}

// LittleMonkey 结构体，嵌入 Monkey（继承）
type LittleMonkey struct {
	Monkey
}

func main() {
	lm := LittleMonkey{Monkey{Name: "小悟空"}}
	lm.climbing()
	fmt.Println(lm.Name)
	lm.Flying() //BirdAble LittleMonkey

	//&lm是BirdAble的实现类，
	//lm不是BirdAble的实现类，
	//所以，此处要传递&lm
	SeeBlueSkyAndCloud(&lm)
	lm.Swimming()
}
