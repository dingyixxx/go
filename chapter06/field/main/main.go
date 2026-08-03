package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}

func main() {
	//struct的属性，如果是 切片 or map，默认值是nil，即，没有分配空间

	var p1 Person
	fmt.Println(p1) //{ 0 [0 0 0 0 0] <nil> [] map[]}
	p1.slice = make([]int, 1)

	p1.map1 = make(map[string]string)

	if p1.ptr == nil {
		fmt.Println("p1.ptr==nil")
	}
	if p1.slice == nil {
		fmt.Println("p1.slice==nil")
	}
	if p1.map1 == nil {
		fmt.Println("p1.map1==nil")
	}

	//p1.ptr==nil
	//p1.slice==nil
	//p1.map1==nil

	//不make或者make的len小于1，直接改变下标0的值，会报错
	//p1.slice[0] = 8

	//append 对 nil 切片是安全的，它会自动分配底层数组。
	//p1.slice = append(p1.slice, 2, 4, 6)

	//p1.map1["color"] = "red"

	fmt.Println(p1.slice)

}
