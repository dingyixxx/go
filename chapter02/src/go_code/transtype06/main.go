package main

import (
	"fmt"
	"strconv"
)

func main() {
	//  string 转 基本数据类型
	var bStr = "true"
	var numStr = "-321"
	var floatStr = "1.2345678"
	var uintStr = "999"
	parseBool, _ := strconv.ParseBool(bStr)     //惊喜地发现goland也可以用.var快捷键
	fmt.Printf("%T,%t\n", parseBool, parseBool) //输入printf，按tab
	i, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Printf("%T,%d\n", i, i)
	parseFloat, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%T,%f\n", parseFloat, parseFloat)
	parseUint, _ := strconv.ParseUint(uintStr, 10, 64)
	fmt.Printf("%T,%d\n", parseUint, parseUint)

	var hello = "hello"
	parseInt, _ := strconv.ParseInt(hello, 10, 64)
	fmt.Printf("%T,%d", parseInt, parseInt)

	//

}
