package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n",
		stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	//业务判断
	stu.Score = score
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

func main() {
	pupil := &Pupil{} //也可以直接在花括号里面，赋值
	//pupil.Student.Name = "小学生图图"
	//pupil.Student.Age = 7
	//pupil.testing()
	//pupil.SetScore(77)
	//pupil.Student.ShowInfo()
	pupil.Name = "小学生图图"
	pupil.Age = 7 //可以不用pupil.Student，可以简化
	pupil.testing()
	pupil.SetScore(77)
	pupil.ShowInfo()
	fmt.Println()

	graduate := &Graduate{}
	graduate.Student.Name = "大学生韦东奕"
	graduate.Student.Age = 20
	graduate.testing()
	graduate.SetScore(106)
	graduate.Student.ShowInfo()

}
