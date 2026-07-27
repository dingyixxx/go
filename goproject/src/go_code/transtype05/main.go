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
	var num4 float64 = 23.456
	var b2 bool = true

	str1 := strconv.FormatInt(int64(num3), 10)
	fmt.Printf("%T,%q \n", str1, str1)

	str2 := strconv.FormatFloat(num4, 'f', 10, 64)
	//说明：'f'格式 10:小数点位保留10位 64：表示这个小数是float64
	fmt.Printf("%T,%q \n", str2, str2)

	str3 := strconv.FormatBool(b2)
	fmt.Printf("str3:%T,%q\n", str3, str3)

	var num11 = 119
	fmt.Println(strconv.Itoa(num11))

}
