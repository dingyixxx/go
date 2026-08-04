package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Stu Student

type integer222 int

func main() {
	//	结构体可以进行type重新定义（相当于取别名），
	//	Golang认为是新的数据类型，
	//	但是相互之间可以强转

	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1) //不能直接等,得强转
	fmt.Println(stu1, stu2)

	var i integer222 = 10
	var j int = 20
	j = int(i)
	fmt.Println(i, j)

}
