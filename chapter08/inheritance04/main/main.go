package main

import "fmt"

type A struct {
	Name string
	Age  string
}

type Stu struct {
	A
	int
}

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int
}

func main() {
	//匿名字段是基本数据类型
	var m E
	m.Name = "史瑞克"
	m.Age = 18
	m.int = 99
	fmt.Println(m)
}
