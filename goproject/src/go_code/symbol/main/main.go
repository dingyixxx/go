package main

import "fmt"

func main() {
	fmt.Println(11 / 4)
	fmt.Println(11 % 4)

	var res float32 = 11 / 4
	fmt.Println(res) //2

	var res1 float32 = float32(11) / float32(4)
	fmt.Println(res1) //2.75

	var res3 float32 = 10.0 / 4
	fmt.Println(res3) //2.5

	fmt.Println(10 % 3)   //1
	fmt.Println(-10 % 3)  //-1
	fmt.Println(10 % -3)  //1
	fmt.Println(-10 % -3) //-1

	var i = 10
	i++
	fmt.Printf("i++之后，%d\n", i)
	i--
	fmt.Printf("i--之后，%d\n", i)

	fmt.Println(19 / -5)

	var a int = 19
	a++ //a++只能独立使用，并不能b:=a++
	//++a 也不能这么写

	//if a++ >0{
	//
	//}go不支持

	//var d int=10
	//d=d++ 不能这么写的
	//fmt.Println(d)

	//var d int=10 不能这么写的
	//if d++ >10{
	//	fmt.Println("ok")
	//}
}
