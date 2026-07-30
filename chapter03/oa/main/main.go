package main //main包只能有一个，如果要编译为可执行文件，就需要将包名定义为main

import (
	"fmt"
	sweetie "go_code/chapter03/oa/utils" //起别名之后，原来的包名就用不了了...
)

// 包的本质，是文件夹
// 包名和文件夹名，通常要保持一致
func main() {
	sweetie.SayHello() //包名.函数名
	sweetie.Cal()
	fmt.Println()
	calc1 := sweetie.Calc1(11.11, 71.43, '*')
	fmt.Println(calc1)
	fmt.Println(sweetie.Num1)
}
