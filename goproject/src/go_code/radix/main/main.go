package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int = 10
	fmt.Println(n)

	var n2 int = 0110 //8进制 0开头
	fmt.Println(n2)

	var n3 int = 0x110 //16进制 0x
	fmt.Println(n3)

	var n4 int = 0x21af //也是16进制 0X
	fmt.Println(n4 + 1)
	fmt.Printf("%x\n", n4+1) //21b0

	var num string = "01001" //9
	numb, _ := strconv.ParseInt(num, 2, 64)
	fmt.Println(numb)
}
