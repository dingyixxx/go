package main

import "fmt"

func main() {
	var s string = "我爱那塞北的雪abc1"
	fmt.Println(len(s)) //25
	var runes []rune = []rune(s)
	fmt.Println(len(runes)) //11

}
