package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("rType的值是%v,类型是%T\n", rType, rType)
	//rType的值是main.Student,类型是*reflect.rtype

	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue的值是%v,类型是%T\n", rValue, rValue)
	//rValue的值是{汤姆 20},类型是reflect.Value

	fmt.Printf("rType.Kind()的值是%v,类型是%T\n", rType.Kind(), rType.Kind())
	fmt.Printf("rValue.Kind()的值是%v,类型是%T\n", rValue.Kind(), rValue.Kind())
	//rType.Kind()的值是struct,类型是reflect.Kind
	//rValue.Kind()的值是struct,类型是reflect.Kind

	//下面我们将rValue转成interface{}
	iV := rValue.Interface()
	fmt.Printf("iV的值是%v,类型是%T\n", iV, iV)
	//iV的值是{汤姆 20},类型是main.Student

	//先判断能不能转成这个类型
	//iV.(type)只能和switch一起用,不能单独用
	switch iV.(type) {
	case Student:
		fmt.Println("类型是Student") //类型是Student
	}

	//将interface{}通过断言,重新转回需要的类型
	//先判断能不能转成这个类型
	res, ok := iV.(Student)
	if ok == false {
		fmt.Println("类型强转Student失败") //类型是Student
	}
	fmt.Printf("res的值是%v,类型是%T\n", res, res)
	//res的值是{汤姆 20},类型是main.Student

}

type Student struct {
	Name string
	Age  int
}

func main() {
	//结构体\interface{}\reflect.Value的转换
	var stu = Student{
		Name: "汤姆",
		Age:  20,
	}
	reflectTest01(stu)
}
