package main

import "fmt"

func main() {
	//	基本数据类型的默认值
	//	在go中，当程序没有赋值时，就会保留默认值，在go中，默认值又叫零值
	//	整型 默认值 0
	//	浮点型 默认值 0
	//	字符串 默认值 ""
	//	布尔类型 默认值 false

	//var a int
	//var b float32
	//var c float64
	//var isReal bool
	//var d string
	////格式化输出 %f 浮点数（小数） %v 任意值的默认格式
	////a=0,b=0.000000,c=0.000000,isReal=false,d=
	//fmt.Printf("a=%d,b=%f,c=%f,isReal=%v,d=%v", a, b, c, isReal, d)

	//	golang中，不同类型的变量之间赋值时需要"显示转换"，即使，低精度向高精度转换的时候，也不能自动转换
	//	表达式T(v)将值v转换为类型T
	//	T就是数据类型，比如，int32，int64，float32等等
	//	v：就是需要转换的变量
	//	案例演示：
	var i int = 42
	var f = float32(i)
	var u uint8 = uint8(f)
	fmt.Println(i, f, u)
	fmt.Printf("%d,%f,%v", i, f, u) //格式化打印，第一个字符串不需要用逗号分割，后续填充的值用逗号分割

}
