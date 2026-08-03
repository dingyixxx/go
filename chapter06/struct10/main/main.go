package main

type Student struct {
	Name string
	Age  int
}

type Stu Student

func main() {
	//	结构体可以进行type重新定义（相当于取别名），
	//	Golang认为是新的数据类型，
	//	但是相互之间可以强转

}
