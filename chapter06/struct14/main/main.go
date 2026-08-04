package main

import (
	"fmt"
	"go_code/chapter06/struct14/utils"
)

func main() {

	//方法的访问范围控制的规则，和函数一样。
	//方法名首字母小写，只能在本包中访问。
	//方法首字母大写，可以在本包和其他包中访问。
	fmt.Println()
	var slice1 = utils.SliceType{}
	slice1.BiggerSpeak(3)
}
