package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := "hello"
	for i := 0; i < 700000; i++ {
		str += strconv.Itoa(i)
	}
}

func main() {
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行了%d秒", end-start)
}
