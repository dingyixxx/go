package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//初始化Person
	//方式一：返回的是值
	var person Person

	//方式二：返回的是值
	var person1 = Person{}

	//方式三：返回的是值（这种方式，用得最多，获得值，而不是指针）
	var person2 = Person{
		Name: "Mary",
		Age:  30,
	} //光标置于花括号之后，直接tab，就可以输入值了，这是一个快捷输入方法

	fmt.Println("person-", person)   //person- { 0}
	fmt.Println("person1-", person1) //person1- { 0}
	fmt.Println("person2-", person2) //person2- {Mary 30}
	(&person2).Name = "May"
	fmt.Println("person2-", person2) //person2- {May 30}

	//	下面两种方式，返回的是结构体指针
	var person3 *Person = new(Person) //方式四：返回的是结构体指针
	(*person3).Name = "王大锤"           //此处，用的是*person3，含义为“该指针指向的对象”
	//person3-地址是：0x1400000e060
	fmt.Printf("person3-地址是：%p\n", person3)

	//person3-值是：&{王大锤 0}
	fmt.Printf("person3-值是：%v\n", person3)
	fmt.Printf("person3-值是：%v\n", *person3)

	person3.Age = 89 //但是go做了简化，也支持 结构体指针.字段名=xxx
	//person3-值是：&{王大锤 89}
	fmt.Printf("person3-值是：%v\n", person3)

	//方式五：返回的是结构体指针
	var person5 *Person = &(Person{})
	//var person5 *Person=&Person{}
	person5.Age = 188
	fmt.Println("person5-", person5) //person5- &{ 188}

	//也可以直接在花括号里面赋值
	var person6 = &(Person{
		Name: "老周",
		Age:  60,
	})
	fmt.Println("*person6-", *person6)

}
