package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2 //不用:=
	sub = n1 - n2 //不用:=
	return        //return不写不行，return后面也不用跟着sum或者sub
}

// go支持可变参数
// 支持0到多个参数
func sum(args ...int) (res int) {
	res = 0
	return
}

func sum1(n1 int, args ...int) (res int) {
	res = n1
	for i := 0; i < len(args); i++ {
		res += args[i] //取出args切片的第一个值
	}
	return
}

// 如果n1和n2数据类型一致的情况下，前者的数据类型可以省略
func add(n1, n2 float32) float64 {
	fmt.Printf("%T\n", n1)
	n1d := decimal.NewFromFloat32(n1)
	f, _ := n1d.Add(decimal.NewFromFloat32(n2)).Float64()
	return f
}

// 非可变参数，必须只能写在第一位，否则会编译不过
//func sum2(args ...int,n1 int) (res int) {
//	res = n1
//	for i := 0; i < len(args); i++ {
//		res += args[i] //取出args切片的第一个值
//	}
//	return
//}

//说明：
//args是slice切片，通过args[index]可以访问到各个值

func main() {
	//	对返回值命名
	//a, b := getSumAndSub(20, 9)
	//fmt.Printf("a=%d,b=%d", a, b)

	_, b := getSumAndSub(20, 9)
	fmt.Printf("b=%d\n", b)

	res1 := sum1(10)               //10
	res2 := sum1(10, 4, 5, 6, 7)   //32
	res3 := sum1(2, 1, 3, 6, 4, 5) //21
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	res := add(1.11, 2.559876723)
	fmt.Printf("res的值为：%f", res)
}
