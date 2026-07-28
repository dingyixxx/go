package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv.atoi和strconv.itoa是比较常用的两个函数
	var num int = 456
	//letter := strconv.Itoa(num)
	letter := strconv.FormatInt(int64(num), 10)
	fmt.Println(letter)

	var str1 string = "123"
	//num2, _ := strconv.Atoi(str1)
	num2, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Println(num2)

}
