package main

import "fmt"

func main() {
	//	value如果也是map,那还要继续make
	studentInfo := make(map[string]map[string]string)
	studentInfo["stu01"] = make(map[string]string)
	studentInfo["stu01"]["name"] = "dingyi"
	studentInfo["stu01"]["age"] = "35"
	studentInfo["stu01"]["email"] = "dingyixxx@126.com"
	studentInfo["stu02"] = make(map[string]string)
	studentInfo["stu02"]["name"] = "dachui"
	studentInfo["stu02"]["age"] = "44"
	studentInfo["stu02"]["email"] = "531557632@qq.com"
	test01(studentInfo["stu02"])
	fmt.Println(studentInfo)
}

func test01(m2 map[string]string) {
	m2["name"] = "修改后的名字"
	m2["age"] = "修改后的年龄"
	m2["email"] = "修改后的邮箱"
}
