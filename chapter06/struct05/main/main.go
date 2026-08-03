package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "张三",
		Age:  5,
	}
	person2 := &person //person2是指针

	//person的地址是：0x1400000e018
	//person2的值是：0x1400000e018
	//person2本身的地址是：0x14000064038
	fmt.Printf("person的地址是：%p\n", &person)
	fmt.Printf("person2的值是：%p\n", person2)
	fmt.Printf("person2本身的地址是：%p\n", &person2)

}
