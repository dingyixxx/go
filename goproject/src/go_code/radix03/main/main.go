package main

import "fmt"

func main() {
	var num = 56
	fmt.Printf("%b\n", num) //111000 8+16+32=56

	//Goland格式化代码的快捷键
	//Option + Command + L
	//大写的L，所以你需要按住左边大写的上箭头
	var num1 = 156
	fmt.Printf("%o\n", num1) //%o不带0前缀 234 4+3*8+2*8*8=156
	fmt.Printf("%O\n", num1) //%0带0前缀 0o234

	var num2 = 356
	fmt.Printf("%x\n", num2) //164 4+6*16+256

	//123转成二进制
	//678转成八进制
	//8912转成十六进制
	var n1 int = 123
	var n2 int = 678
	var n3 int = 8912
	fmt.Printf("%b\n", n1) //1111011 1+2+8+16+32+64=123
	fmt.Printf("%O\n", n2) //0o1246 6+4*8+2*8*8+8*8*8=678
	fmt.Printf("%x\n", n3) //22d0 16*13+2*16*16+2*16*16*16=8912

}
