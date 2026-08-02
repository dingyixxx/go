package main

import "fmt"

func main() {
	//	map的使用, 有三种方式

	//	方式一
	//声明+make+赋值
	var a map[string]int
	a = make(map[string]int, 0) //即使写0, 也可以存东西的
	a["a"] = 1
	fmt.Println(a)

	//不写capacity, 也可以
	var b map[string]int
	b = make(map[string]int)
	b["a"] = 1
	fmt.Println(b)

	//	方式二
	//声明,直接make
	//再赋值
	var cities map[int]string = make(map[int]string)
	cities[21] = "上海"
	cities[10] = "北京"
	fmt.Println(cities)

	//	方式三
	//	声明,直接赋值
	//就不需要make了
	var goods map[string]string = map[string]string{
		"no1": "beijing",
		"no2": "shanghai",
		"no3": "guangzhou", //此处这个逗号, 没有不行

	}
	fmt.Println(goods)

}
