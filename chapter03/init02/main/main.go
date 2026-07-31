package main

import (
	"fmt"
	"go_code/chapter03/init02/utils"
)

// 神奇，var和func都有变量提升
var num = getNum()

func getNum() int {
	fmt.Println("main包...getNum执行中...")
	return 119
}

func init() {
	fmt.Println("main包...init...")
}

func main() {
	fmt.Println("main包...main...")
	fmt.Printf("main包...定义的变量num：%d\n", num)
	fmt.Printf("utils.UtilName的变量是：%v,utils.UtilAge的变量是：%v", utils.UtilName, utils.UtilAge)
}
