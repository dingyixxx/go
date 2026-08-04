package main

import "fmt"

func main() {
	var p *int64
	fmt.Println(*p) //panic: runtime error: invalid memory address or nil pointer dereference
}
