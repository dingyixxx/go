package main

type Goods struct {
	Name  string
	Price int
}

type Book struct {
	Goods
	Writer string
}

func main() {
	//	无extends
	//	封装，encapsulation，把抽象出的“字段和对字段的操作”，封装在一起，数据被保护在内部，程序的其他包，只有通过被授权的操作（方法），才能对字段进行操作。

	//	继承 匿名结构体
}
