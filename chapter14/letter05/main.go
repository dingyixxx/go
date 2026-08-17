package main

import "fmt"

func main() {
	var str string = "我爱那塞北的雪"
	for i, v := range str {
		fmt.Printf("i=%v,v=%c\n", i, v)
	}
}
