package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Age = 10
	p1.Name = "小名"
	var p2 *Person = &p1
	fmt.Println((*p2).Age) //必须加括号，
	// 否则，会认为是对于“p2.Age”解指针，但p2.Age并不是一个指针

}
