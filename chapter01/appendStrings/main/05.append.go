package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "hello"
	b := 3
	fmt.Println(a + strconv.Itoa(b)) //a和b不能直接相加
}
