package main

import "fmt"

func main() {
	var data *byte
	var in interface{}
	fmt.Println(data, data == nil) //prints: <nil> true
	fmt.Println(in, in == nil)     //prints: <nil> true
	in = data
	fmt.Println(in, in == nil) //prints: <nil> false
	//'data' is 'nil', but 'in' is not 'nil'

	//in == nil 要求 type 和 value 同时为 nil 才返回 true。in = data 后，type 变成了 *byte（非 nil），所以 in != nil

	//普通类型（如 *byte）只有一个值，值是 nil 就是 nil。
	//interface 类型有 type + value 两个字段，必须两者都为 nil 才算 nil。
	//给 interface 赋值了一个 nil 的具体类型后，type 字段被填充，就不再是 nil 了。

	//interface{}判断是否为nil的陷阱
}

func main2() {

	doit := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		}
		return result
	}

	if res := doit(-1); res != nil {
		fmt.Println("good result:", res) //prints: good result: <nil>
		//'res' is not 'nil', but its value is 'nil'
	}
}
