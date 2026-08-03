package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["stu01"] = "小梅"
	m["stu02"] = "小李"
	m["stu03"] = "小王"

	delete(m, "stu03")
	delete(m, "stu-999") //也不报错

	fmt.Println(m)

	itsName, res := m["stu02"]
	fmt.Println(itsName, res)

	itsName1, res1 := m["stu333"]
	fmt.Println(itsName1, res1) //false
}
