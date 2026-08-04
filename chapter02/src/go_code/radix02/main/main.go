package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var i int = 5 //101
	//n := fmt.Sprintf("%b", i)
	//fmt.Println(i)
	//fmt.Println(n)
	//fmt.Printf("%T,%v", n, n) //string,101

	var num1 int = 0b110001100
	var num2 int = 02456
	var num3 int = 0xa45
	fmt.Printf("%d\n", num1) //4+8+128+256=396 0b 0b 0b
	fmt.Printf("%d\n", num2) //6+5*8+4*8*8+2*8*8*8=1326
	fmt.Printf("%d\n", num3) //5+4*16+10*16*16=2629

	var str1 string = "110001100"
	//var str2 string = "02456"
	var str22 string = "2456" //如果不带前缀0，base就传入8，带的话base传0
	//var str3 string = "0xa45"
	var str33 string = "a45" //如果不带前缀0x，base就传入16，带的话base传0

	num1d, _ := strconv.ParseInt(str1, 2, 64)
	println(num1d)

	//num2d, _ := strconv.ParseInt(str2, 0, 64)
	num2d, _ := strconv.ParseInt(str22, 8, 64)
	println(num2d)

	//num3d, _ := strconv.ParseInt(str3, 0, 64)
	num3d, _ := strconv.ParseInt(str33, 16, 64)
	println(num3d)

}
