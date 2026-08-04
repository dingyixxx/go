package main

import (
	"fmt"
	"go_code/chapter07/factory/model"
)

func main() {
	//var s = model.Student{
	//	Name:  "张三",
	//	Score: 78.9,
	//}
	//fmt.Println(s)

	//var s = model.person{
	//	Name:  "张三",
	//	Score: 78.9,
	//}
	//fmt.Println(s) 小写的person，不能跨包直接用

	person := model.NewPerson("tom", 88.89)
	fmt.Println(*person)
	fmt.Printf("%T\n", person) //*model.person
	fmt.Println(person.Name, person.GetScore())

}
