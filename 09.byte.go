package main

import "fmt"

func main() {
	//	golang没有专门的字符类型，如果要存储单个字符（字母），一般使用byte来保存
	//	go的字符串是由单个字节连接起来的。传统的字符串是由字符组成的。

	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	//	当我们直接输出byte值，就是输出了对应的字符的码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)

	fmt.Printf("c1=%c c2=%c", c1, c2)

	//var c3 byte = '北' //byte的取值范围是0-255，北是不可以的 cannot use '北' (untyped rune constant 21271) as byte value in variable declaration (overflows)
	//ascii表([0-1,a-z,A-Z])里的用byte来存放，否则，如果码值大于255，就用int来存放
	//fmt.Printf("c3=%c", c3)

	var c3 int = '北'                       //这个是可以的
	fmt.Printf("c3=%c 对应码值=%d \n", c3, c3) //格式化输出的就不是码值了，而是
	fmt.Println(c3)                        //21271

	var c4 byte = 97
	fmt.Printf("c4的格式化打印结果为：%c \n", c4) //c4的格式化打印结果为：a

	var c5 byte = '\t'
	fmt.Printf("c5=%c", c5)
	fmt.Println(c5) //9

	//	百度搜索“utf-8 编码表 在线” https://www.sojson.com/unicode.html 查看中文的unicode码
	//unicode也是兼容utf-8的
	//在golang中，中文都是通过utf-8编码的
	//	utf-8是unicode的一种实现
	//	ascii码标识中文乏力，后面出现了utf-8，兼容ascii码
	//	英文 1个字节 汉字 3个字节

	//	国22269

	fmt.Printf("%c \n", 120)

	//	字符类型是可以进行运算的，相当于一个整数，因为它都有对应的unicode码
	var res1 int = '国'
	var res2 int = '破'
	var res3 int = '山'
	var res4 int = '河'
	var res5 int = '在'
	var res6 int = '城'
	var res7 int = '春'
	var res8 int = '草'
	var res9 int = '木'
	var res10 int = '深'

	var total = res1 + res2 + res3 + res4 + res5 + res6 + res7 + res8 + res9 + res10
	fmt.Print(total, '\n') //should be 263634

	var total1 = 10 + 'a'
	fmt.Println(total1) //107

}
