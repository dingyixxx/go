package main

import "fmt"

type Usb interface {
	Say()
}

// 方法都传入指针，而不是值
type Stu struct{}

func (this *Stu) Say() { // ⚠️ 指针接收者
	fmt.Println("Say()")
}

func main() {
	var stu Stu = Stu{}
	//var u Usb = stu // ❌ 编译报错！
	var u Usb = &stu
	u.Say()
	fmt.Println("here", u)
}
