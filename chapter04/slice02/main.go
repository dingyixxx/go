package main

import (
	"fmt"
)

func main() {
	exp := "3+2*6-2/a+b*c"
	//数栈 符号栈（是符号时，判断其和栈顶符号的优先级，进而选择“弹两次数栈+弹一次符号栈+结果入数栈”or“入符号栈”）
	var expSlice []byte = []byte(exp)
	fmt.Printf("expSlice的类型是:%T\n", expSlice)
	ch := expSlice[3:4]
	fmt.Printf("ch的类型是:%T\n", ch)
	fmt.Println(ch[0] == '*')
	for i := 0; i < len(exp); i++ {
		var num byte = exp[i]
		var letter string = string(exp[i : i+1]) //go的字符串，是"字节"数组，不是"字符"数组

		fmt.Printf("ascii码是:%v,字母是:%v\n", num, letter)
	}
	//byte == uint8
	var numb byte = 1
	fmt.Printf("numb的类型是:%T\n", numb) //小名byte，大名uint8，无符号位
	//rune == int32
	var numr rune = '锤' //小名rune，大名int32，有符号位
	fmt.Printf("numr的类型是:%T\n", numr)
	//为什么rune不是uint32而是int32，中文还能为负数吗？
	//不能为负数，就是这么规定的

	//为什么java里的字符串是“字符”数组，golang的字符串是“字节”数组
	//java选择了“固定宽度”，用“空间”换取“索引方便”
	//golang选择了“变长”，用“需要转换成[]rune”换取“内存的使用效率”
	var str2 string = "我爱那塞北的雪"

	fmt.Printf("len(str2)长度%d\n", len(str2))
	//len(str2)长度21

	fmt.Printf("len([]rune(str2))长度%d", len([]rune(str2)))
	//len([]rune(str2))长度7

	//打印str2试试
	fmt.Println("str2[0]", str2[0])
	fmt.Println("string(str2[0:0])", string(str2[0:1])) //string(str2[0:0]) �

	//存英文的话，java浪费一半儿空间

}
