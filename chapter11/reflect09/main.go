package main

import (
	"fmt"
	"reflect"
)

// 编写一个Cal结构体，有两个字段Num1，和Num2。
// 方法GetSub(name string)
// 使用返回遍历Cal结构体所有的字段信息。
// 使用反射机制完成对GetSub的调用，输出形式为 "tom 完成了减法运行，8 - 3 = 5"

type Cal struct {
	Num1 int
	Num2 int
}

func (cal *Cal) GetSub(name string) (res int) {
	return cal.Num1 + cal.Num2
}

func IamReflectingYouDearFunc(obj interface{}, args ...interface{}) {
	rVal := reflect.ValueOf(obj)

	// 遍历结构体字段
	rType := rVal.Elem().Type()
	for i := 0; i < rType.NumField(); i++ {
		fmt.Printf("字段%d: %s = %v\n", i, rType.Field(i).Name, rVal.Elem().Field(i).Interface())
	}

	// 设置字段值
	rVal.Elem().FieldByName("Num1").SetInt(int64(args[0].(int)))
	rVal.Elem().FieldByName("Num2").SetInt(int64(args[1].(int)))

	// 反射调用 GetSub
	method := rVal.MethodByName("GetSub")
	inArgs := []reflect.Value{reflect.ValueOf(args[2].(string))}
	results := method.Call(inArgs)

	fmt.Printf("%s 完成了减法运行，%d - %d = %d\n",
		args[2], args[0], args[1], results[0].Int())

}

func main() {

	cal := &Cal{
		Num1: 8,
		Num2: 3,
	}

	IamReflectingYouDearFunc(cal, 8, 3, "tom")

}
