package main

import (
	"fmt"
	"strconv"
)

func add() func(int) int {
	n := 10
	return func(step int) int {
		n += step
		return n
	}
}

func main() {
	f := add()
	fmt.Println(f(1)) //把n理解为javascript的window.n
	fmt.Println(f(2)) //闭包=内存泄漏
	fmt.Println(f(3))

	num := 1
	fmt.Printf("%v\n", num)

	fmt.Printf("%v\n", string(36)) //36对应的字符
	fmt.Printf("%v\n", string(37))
	fmt.Printf("%v\n", string(38))

	fmt.Println(strconv.Itoa(1))
	fmt.Println(strconv.Itoa(2))
	fmt.Println(strconv.Itoa(3))

}
