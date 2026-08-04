package main

import "fmt"

func main() {
	var a *int = nil
	var b *int = nil
	var c int = 1
	var d *int = &c
	fmt.Println(a == b)
	fmt.Println(a == d)

	//fmt.Println(nil == nil)
	//invalid operation: nil == nil (operator == not defined on untyped nil)
}
