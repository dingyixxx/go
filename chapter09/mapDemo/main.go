package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["a"] = 1

	m2 := m1     // 赋值
	m2["a"] = 99 // 修改 m2

	fmt.Println(m1["a"]) // 99  ← m1 也变了！说明共享同一块数据

	//map是引用类型

}
