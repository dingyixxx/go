package main

import "fmt"

type Stu struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	students := make(map[string]Stu, 0)
	students["no1"] = Stu{"tom", 10, 99.4} //花括号，不是圆括号
	students["no2"] = Stu{"mary", 18, 18.6}
	students["no3"] = Stu{"baby", 20, 77.7}
	for k, v := range students {
		fmt.Printf("编号为%v的学生姓名是%v\n", k, v.Name)
		fmt.Printf("编号为%v的学生年龄是%v\n", k, v.Age)
		fmt.Printf("编号为%v的学生分数是%v\n", k, v.Score)
		fmt.Println()
	}

}
