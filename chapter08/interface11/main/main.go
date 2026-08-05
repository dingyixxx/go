package main

import (
	"fmt"
	"sort"
)

// 声明Hero
type Hero struct {
	Name string
	Age  int
}

// 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 实现Interface（这个接口的名字比较特殊，这个接口的名字就叫Interface）的Len Less Swap三个接口
func (hs *HeroSlice) Len() int {
	return len(*hs)
}

func (hs *HeroSlice) Less(i, j int) bool {
	//return (*hs)[i].Age < (*hs)[j].Age
	return (*hs)[i].Name < (*hs)[j].Name

}

func (hs *HeroSlice) Swap(i, j int) {
	//temp := (*hs)[i]
	//(*hs)[i] = (*hs)[j]
	//(*hs)[j] = temp
	(*hs)[i], (*hs)[j] = (*hs)[j], (*hs)[i] //tuple
}
func main() {
	var slice = HeroSlice([]Hero{{
		Name: "a22大锤",
		Age:  22,
	},
		{
			Name: "c18大锤",
			Age:  18,
		}, {
			Name: "b7大锤",
			Age:  7,
		}, {
			Name: "d112大锤",
			Age:  112,
		}, {
			Name: "e33大锤",
			Age:  33,
		}, {
			Name: "g12锤",
			Age:  12,
		}, {
			Name: "f5大锤",
			Age:  5,
		}})
	sort.Sort(&slice)
	fmt.Println(slice)
}
