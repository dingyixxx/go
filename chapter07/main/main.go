package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student struct {
	StudentNo string
	Score     int
}

func (p Person) test() {
	p.Age++
	fmt.Printf("Person名字是：%v，年龄是：%v\n", p.Name, p.Age)
}

func hanshuP(p Person) {
	p.Age++
	fmt.Printf("Person名字是：%v，年龄是：%v\n", p.Name, p.Age)
}

// 方法又可以重名了
func (s *Student) test() {
	s.Score++
	s.Score++
	s.Score++
	s.Score++

	fmt.Printf("Student学号是：%v，分数是：%v\n", s.StudentNo, s.Score)
}

func hanshuS(s *Student) {
	s.Score++
	s.Score++
	s.Score++
	s.Score++

	fmt.Printf("Student学号是：%v，分数是：%v\n", s.StudentNo, s.Score)
}

func main() {
	var p Person = Person{
		Name: "大锤",
		Age:  35,
	}
	(&p).test()
	//hanshuP(&p)这么写是错的
	//hanshuP(p)这么写才是对的
	fmt.Println("main里面打印的-", p)

	fmt.Println()
	var s Student = Student{
		StudentNo: "No1001",
		Score:     96,
	}
	s.test()
	//hanshuS(s) //这么写是错的
	//hanshuS(&s)这么写才是对的
	fmt.Println("main里面打印的-", s)

}
