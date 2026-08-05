package main

import "fmt"

type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
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
	var UsbArray = [3]Usb{
		Computer{
			Name: "Dell xps计算器",
		}, Camera{
			Mingzi: "尼康",
		}, Phone{
			Xinghao: "vivo",
		},
	}
	for _, value := range UsbArray {
		value.Start()
		value.Stop()
		if res, err := value.(Phone); err == true {
			res.Call119()
		}
	}
}
