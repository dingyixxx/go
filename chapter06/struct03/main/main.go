package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	xiaobai := Student{"小白", 35}
	ming := xiaobai
	ming.Name = "小明"
	fmt.Println(xiaobai)

	hong := &xiaobai
	hong.Name = "变为小红"
	fmt.Println(xiaobai)

}
