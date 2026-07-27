package main

import "fmt"

func main() {
	num := 999
	ptr := &num //短变量声明，无须指定类型
	fmt.Printf("%v\n", ptr)

	num1 := 300
	//var ptr1 *float32=&num1 //编译不通过
	var ptr1 = &num1 //编译不通过
	fmt.Printf("ptr1=%d\n", ptr1)

	var a int = 300
	var b int = 400
	var ptr2 *int = &a
	fmt.Printf("&a=%d\n", &a)
	fmt.Printf("ptr2=%d\n", ptr2)
	fmt.Printf("*ptr2=%d\n", *ptr2)

	*ptr2 = 100 //a 100
	ptr2 = &b
	*ptr2 = 200 //b 200
	fmt.Printf("&b=%d\n", &b)
	fmt.Printf("ptr2=%d\n", ptr2)
	fmt.Printf("a=%d,b=%d,*ptr2=%d\n", a, b, *ptr2)

	//&a=1374389600432
	//ptr2=1374389600432
	//*ptr2=300
	//&b=1374389600440
	//ptr2=1374389600440
	//a=100,b=200,*ptr2=200

}
