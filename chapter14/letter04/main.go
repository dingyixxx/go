package main

import (
	"fmt"
)

func main() {
	//单个字符可能占用多个rune，行吧
	data := "é"
	fmt.Println(len(data)) //prints: 3
	var runes []rune = []rune(data)
	fmt.Println(len(runes)) //2
}
