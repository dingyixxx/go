package main

import "fmt"

// 接口的继承,和Java一样的

// 也是Duck Typing
type SayI interface {
	Say()
}

type HelloI interface {
	Hello()
}
type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}
func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
}

func main() {
	//Monster实现了SayI和HelloI
	m := Monster{}
	var s SayI = m
	var h HelloI = m
	s.Say()
	h.Hello()
}
