package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "01011"
	num, _ := strconv.ParseInt(str, 2, 8)
	fmt.Println(num)

	//	第二种方式
	var num3 = 99
	//var num4 float64=23.456
	//var b2 bool=true

	str1 := strconv.FormatInt(int64(num3), 10)
	fmt.Printf("%T,%q \n", str1, str1)

}
