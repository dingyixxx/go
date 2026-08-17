package main

import "fmt"

func main() {
	// 类型断言
	//var data interface{} = "great"
	//if data, ok := data.(int); ok {
	//	fmt.Println("[is an int] value =>", data)
	//} else {
	//	fmt.Println("[not an int] value =>", data)
	//	//prints: [not an int] value => 0 (not "great")
	//}

	var data interface{} = "great"
	if res, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not an int] value =>", data)
		//prints: [not an int] value => great (as expected)
	}

}
