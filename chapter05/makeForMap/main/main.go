package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 1)
	a["11"] = "value1"
	a["2"] = "value2"      //内存空间即使分配了1,也不会报错,且也可以存下来
	a["33"] = "value1-new" //不会报错,会覆盖
	a["4"] = "value3"
	a["6"] = "value4"
	a["11"] = "value5"

	fmt.Println(a)
	//	每次打印出来, 没有测试出来无序

}
