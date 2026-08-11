package main

import (
	"fmt"
	"reflect"
)

// 通过反射, 修改num int的值
func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal.Kind()是%v\n", rVal.Kind())
	//rVal.Kind()是ptr 如果传入的是int，则此处显示为int

	//rVal.SetInt(20)//
	rVal.Elem().SetInt(33)
	//Elem返回v持有的指针Elem指向的值的Value封装

	//类似于这么处理
	var numA int = 99
	var numAPtr = &numA
	numBPtr := numAPtr
	*numBPtr = 1000
	fmt.Printf("numA-%v\n", numA)
	fmt.Printf("*numAPtr-%v\n", *numAPtr)

	//	panic: reflect: reflect.Value.SetInt using unaddressable value
	//	编译的时候不报错，但是运行时报了这个错
}

func main() {
	var num int = 10
	reflect01(&num) //如果需要改变，则必须传“引用”
	fmt.Println("num=", num)
}
