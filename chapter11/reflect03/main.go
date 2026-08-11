package main

import (
	"fmt"
	"reflect"
)

func main() {
	//给你一个变量 var v float64 = 1.2，请使用反射来得到它的 reflect.Value，然后获取对应的 Type、Kind 和值，并将 reflect.Value 转换成 interface{}，再将 interface{} 转换成 float64。

	var v float64 = 1.2
	reflectTest001(v)

}

func reflectTest001(b interface{}) {
	rVal := reflect.ValueOf(b)
	iv := rVal.Interface()
	f := iv.(float64)
	fmt.Printf("f的类型是%T，值是%v", f, f)
}
