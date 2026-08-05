package main

import "fmt"

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	Test03()
}

type Stu struct{}

func (stu Stu) Test01() {}
func (stu Stu) Test02() {}
func (stu Stu) Test03() {}

func main() {
	stu := Stu{}
	var a AInterface = stu
	var b BInterface = stu
	fmt.Println("ok~", a, b)
}
