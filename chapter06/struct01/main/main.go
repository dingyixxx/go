package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Version string
}

func main() {
	//golang不能重写equals，严格用==比较每一个属性
	m := map[Person]string{
		{"丁一", 35, "1.1"}: "工程师",
		{"张三", 28, "1.1"}: "设计师",
	}

	fmt.Println(m[Person{"丁一", 35, "1.2"}])
	fmt.Println("------")
	fmt.Println(m[Person{"丁一", 35, "1.1"}])

}
