package main

import "fmt"

func main() {
	var s string = "😊123" //颜文字，4个字节
	fmt.Println(len(s))   //7
	var runes []rune = []rune(s)
	fmt.Println(len(runes)) //4
}
