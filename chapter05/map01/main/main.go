package main

import "fmt"

func main() {
	//map的声明, 声明一个map是不会分配内存的,初始化需要make,分配内存后才能赋值和使用
	var map01 map[string]map[int]int
	fmt.Println(map01)

	var slice []string
	fmt.Println(slice)

	var a map[string]string
	//没有空间,使用map前需要先make
	//a["key1"] = "value1" //panic: assignment to entry in nil map
	fmt.Println(a)
}
