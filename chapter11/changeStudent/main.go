package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  float64
}

// 通过反射, 修改stu Student的值
func main() {
	stu := Student{
		Name: "张三",
		Age:  32,
	}
	reflectTest02(&stu)
	fmt.Println(stu)
}

func reflectTest02(obj interface{}) {
	rVal := reflect.ValueOf(obj)
	rVal.Elem().FieldByName("Name").SetString("李四")
	rVal.Elem().FieldByName("Age").SetFloat(35.2)

	//panic: reflect: call of reflect.Value.SetBool on float64 Value
	//rVal.Elem().FieldByName("Age").SetBool(true)

}
