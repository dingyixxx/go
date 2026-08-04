package main

import (
	"fmt"
	"go_code/chapter02/src/go_code/valuetype/model"
)

func main() {
	//	值类型和引用类型
	//	1）值类型
	//	基本数据类型int系列，float系列，bool，string、数组和结构体struct
	//	变量直接存储值，内存通常在栈中分配
	//	逃逸分析，有可能一个值类型，也分配在堆区

	//var 3num int
	//fmt.Println(num)

	var num = 10
	fmt.Println(num)
	var Num = 222
	fmt.Println(Num)

	var break1 = 44
	fmt.Println(break1) //保留关键字不能作为变量名，例如，chan fallthrough goto

	//var _ int=989 //Cannot use '_' as a value
	//println(_)

	//	2）引用类型
	//	指针，slice切片，map，管道channel、interface
	//	变量存储的是一个地址，这个地址对应的空间才真正存储数据（值），内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收。
	//	逃逸分析，有可能一个引用类型，也分配在栈区

	//	标识符的命名规范（变量、方法、包）
	//	1）26个英文字母大小写、0-9、_组成
	//	2）数字不可以开头
	//	3）严格区分大小写
	//	4）不能包含空格
	//	5）下划线“_”在go中是一个特殊的标识符，成为“空标识符”。可以代表任何其他的标识符，但是它对应的值会被忽略（比如：忽略某个返回值）。所以仅能被作为占位符使用，不能作为标识符使用。
	//	6）不能以系统关键字作为标识符，比如，break，if等等...

	//	标识符的举例说明
	var int = 44
	fmt.Println(int) //int居然可以...

	var float32 float32 = 44.33
	fmt.Println(float32) //float32居然可以...

	//	go项目文件命名用下划线、全小写，不能用小驼峰或者大驼峰
	//包名尽量和文件夹名字保持一致，文件夹名字package04-包名main-文件名main.go
	//	变量、函数、常量采用驼峰

	//	go语言非常特殊的一个特性
	//如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用（注：可以简单地理解成，首字母大写是公有的，首字母小写诗私有的）
	fmt.Println(model.HeroName)

}
