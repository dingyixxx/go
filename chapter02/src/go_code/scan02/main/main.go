package main

import "fmt"

func main() {
	//	go语言里面的字符串,就是slice of bytes 看到%s,读到k i m,遇到空格,停止,把kim塞到&name 不能用%v
	//	%d是读取一段连续的十进制数字
	//	%f是匹配浮点

	var name string
	var age int
	var sal float64
	var isPass bool

	fmt.Println("请输入你的姓名,年龄,工资和是否通过:")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名是:%v,年龄是:%d,工资是:%v,是否通过:%t", name, age, sal, isPass)

}
