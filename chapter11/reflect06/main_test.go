package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestReflectFunc(t *testing.T) {
	sumFunc := func(v1 int, v2 int) {
		fmt.Println("sumFunc函数执行了->", v2+v1)
	}
	minusFuncAndAppendStr := func(v1 int, v2 int, s string) {
		fmt.Println("minusFuncAndAppendStr函数执行了v2-v1再拼接字符串->", strconv.Itoa(v2-v1)+"-"+s)

	}
	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
		for i := 0; i < n; i++ {
			fmt.Printf("%v,", inValue[i])
		}
		fmt.Println()
	}
	bridge(sumFunc, 11, 222)
	bridge(minusFuncAndAppendStr, 111, 22, "suffix")
}
