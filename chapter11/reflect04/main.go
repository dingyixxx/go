package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var str string = "tom"
	//fs := reflect.ValueOf(str) // ❌ 值类型，不可寻址
	//fs.SetString("jack")
	//fmt.Printf("%v\n", str) // tom

	var str string = "tom"
	fs := reflect.ValueOf(&str) // 传指针
	fs.Elem().SetString("jack") // 解引用后修改
	fmt.Printf("%v\n", str)     // jack
}
