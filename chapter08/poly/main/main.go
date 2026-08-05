package main

import "fmt"

type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
	//Test() //如果接口要求实现Test,那么,如果传入的结构体没有实现该方法,则会报错
}

type Phone struct {
	Xinghao string
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Call119() {
	fmt.Println("报警")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
	Mingzi string
}

// 让Camera实现Usb接口的方法
func (camera Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (camera Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
	Name string
}

// 让Computer实现Usb接口的方法
func (computer Computer) Start() {
	fmt.Println("计算机开始工作...")
}

func (computer Computer) Stop() {
	fmt.Println("计算机停止工作...")
}

func main() {
	//多态数组
	var UsbArray [3]Usb = [3]Usb{
		Computer{
			Name: "Dell xps计算器",
		}, Camera{
			Mingzi: "尼康",
		}, Phone{
			Xinghao: "vivo",
		},
	}
	fmt.Println(UsbArray)

	//多态切片
	var UsbSlice []Usb = []Usb{
		Computer{
			Name: "苹果的mac电脑",
		}, Camera{
			Mingzi: "佳能相机",
		}, Phone{
			Xinghao: "tcl翻盖手机",
		},
	}
	fmt.Println(UsbSlice)

}
