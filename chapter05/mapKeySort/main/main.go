package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的key是无序的，每次遍历key，输出的结果都不一样
	m := make(map[int]int)
	m[12] = 100
	m[11] = 13
	m[4] = 56
	m[8] = 90

	for k, v := range m {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}

	//k=12,v=100
	//k=11,v=13
	//k=4,v=56
	//k=8,v=90

	//k=8,v=90
	//k=12,v=100
	//k=11,v=13
	//k=4,v=56

	//k=11,v=13
	//k=4,v=56
	//k=8,v=90
	//k=12,v=100

	//解决方法是，先把key加到一个切片里，排序key，再输出value
	sliceOfKeys := make([]int, 0)
	for key, _ := range m {
		sliceOfKeys = append(sliceOfKeys, key)
	}
	sort.Ints(sliceOfKeys)
	fmt.Println(len(sliceOfKeys))
	for i := 0; i < len(sliceOfKeys); i++ {
		fmt.Printf("m的key为%v的值为%v\n", sliceOfKeys[i], m[sliceOfKeys[i]])
	}

}
