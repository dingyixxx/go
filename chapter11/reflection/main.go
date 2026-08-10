package main

import (
	"fmt"
	"reflect"
)

func bridge(fn interface{}, args ...interface{}) interface{} {
	//处理
	return nil
}

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("rType的值是%v,类型是%T\n", rType, rType)
	//rType的值是int,类型是*reflect.rtype

	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue的值是%v,类型是%T\n", rValue, rValue)
	//rValue的值是11,类型是reflect.Value

	fmt.Printf("rType.Kind()的值是%v,类型是%T\n", rType.Kind(), rType.Kind())
	fmt.Printf("rValue.Kind()的值是%v,类型是%T\n", rValue.Kind(), rValue.Kind())
	//rType.Kind()的值是int,类型是reflect.Kind
	//rValue.Kind()的值是int,类型是reflect.Kind

	//kind 大类 家用电器 mankind 人类
	//type 微波炉 O型血的人 确实型人

	n2 := 2 + rValue.Int()
	//n2 := 2 + rValue.Float()
	//编译阶段检查不出来, 运行时才报错
	//panic: reflect: call of reflect.Value.Float on int Value

	fmt.Println("n2---", n2)

	//下面我们将rValue转成interface{}
	iV := rValue.Interface()
	fmt.Printf("iV的值是%v,类型是%T\n", iV, iV)
	//iV的值是11,类型是int

	//将interface{}通过断言,重新转回需要的类型
	res := iV.(int)
	fmt.Printf("res的值是%v,类型是%T\n", res, res)
	//res的值是11,类型是int

}

func main() {
	//	1.json tag是反射
	//	2.适配器函数(例如,上面的bridge), 也是反射

	//Type

	//Value

	//基本数据类型\interface{}\reflect.Value的转换
	var num int = 11
	reflectTest01(num)

}
