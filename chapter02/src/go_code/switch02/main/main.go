package main

import "fmt"

func test(b byte) byte {
	return b + 1 - 1
}

func main() {
	var dayCode byte
	fmt.Println("请输入以下字符，a,b,c,d,e,f,g")
	fmt.Scanf("%c", &dayCode)

	//switch 后面 输入 变量、字面量或者函数的返回值，其实都可以
	switch test(dayCode) {
	case 'a':
		fmt.Println("周一，猴子穿新衣")
	case 'b':
		fmt.Println("周二，猴子当小二")
	case 'c':
		fmt.Println("周三，猴子侃大山")
	case 'd':
		fmt.Println("周四，猴子你太放肆")
	case 'e':
		fmt.Println("周五，猴子恰恰舞")
	case 'f':
		fmt.Println("周六，猴子666")
	case 'g':
		fmt.Println("周日，猴子涂油漆")
	default:
		fmt.Println("panic")
	}

}
