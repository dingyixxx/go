package main

import "fmt"

func main() {
	if '*' == 42 {
		fmt.Println("'*' == 42")
	}
	fmt.Println("'A' == 65", 'A' == 65)
	var num int32 = '国'
	fmt.Println(num)
	fmt.Println(string([]rune{num}))

	fmt.Println(22269 == '国') //true

	str2 := "国"
	var byteArray []byte = []byte(str2)
	fmt.Println(byteArray[0])
	fmt.Println(byteArray[1])
	fmt.Println(byteArray[2])
	fmt.Println(str2[0])
	fmt.Println(str2[1])
	fmt.Println(str2[2])
	fmt.Printf("byteArray[0]转二进制是：%b\n", byteArray[0]) //11100101 去掉前缀 1110 为0101
	fmt.Printf("byteArray[1]转二进制是：%b\n", byteArray[1]) //10011011 去掉前缀 10 为011011
	fmt.Printf("byteArray[2]转二进制是：%b\n", byteArray[2]) //10111101 去掉前缀 10 为111101

	var runeArray []rune = []rune(str2)
	fmt.Printf("runeArray[0]转二进制是：%b\n", runeArray[0])
	//101011011111101
	//0101 011011 111101

	fmt.Println(runeArray[0])
	//	字母、符号、数字，无所谓，byte和int互转，不会出错的
	//	如果是“中文”和“颜文字”，那一定一定要注意了，感觉很容易错，一定要转为rune[]

}
