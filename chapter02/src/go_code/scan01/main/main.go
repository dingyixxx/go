package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name) //传的是指针,因为可以更新内容
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)
	fmt.Println("请输入工资:")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试:")
	fmt.Scanln(&isPass)

	fmt.Printf("我的姓名是:%v,我的年龄是:%d,我的工资是:%v,我的考试通过状况是:%t", name, age, sal, isPass)

}
