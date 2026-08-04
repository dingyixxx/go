package main

import "fmt"

type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
	//Test() //如果接口要求实现Test,那么,如果传入的结构体没有实现该方法,则会报错
}

type Phone struct {
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
}

// 让Camera实现Usb接口的方法
func (camera Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (camera Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

// 让Computer实现Usb接口的方法
func (computer Computer) Start() {
	fmt.Println("计算机开始工作...")
}

func (computer Computer) Stop() {
	fmt.Println("计算机停止工作...")
}

// 编写一个方法——working方法
// 接收一个Usb接口类型变量
// 只要是实现了Usb接口
// 所谓“实现”Usb接口，就是指实现了Usb接口的所有方法
func (c Computer) Working(usb Usb) {
	//通过Usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}
func main() {
	//	测试
	//	先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)

}
