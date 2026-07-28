package main

import "fmt"

func test() bool {
	fmt.Println("test...")
	return true
}
func testReturnInt() int {
	return 45
}

func main() {
	//var a int = 313
	//var b int = 2121212
	//fmt.Printf("%b\n", a) //                1 0011 1001
	//fmt.Printf("%b\n", b) //10 0000 0101 1101 1111 1100
	//
	//i := a ^ b
	//fmt.Printf("%b\n", i)
	//10 0000 0101 1101 1111 1100

	//10 0000 0101 1100 1100 0101

	//var a *int = nil
	//var b *int = nil
	//fmt.Println(a == b)   // true（两个同类型 nil 指针相等）
	//fmt.Println(a == nil) // true

	//fmt.Println(nil == nil)

	//var age int = 40
	//if age > 30 && age < 50 {
	//	fmt.Println("ok1")
	//}
	//if age > 30 && age < 40 {
	//	fmt.Println("ok2")
	//}
	//if age < 30 || test() {
	//	fmt.Println("ok3...")
	//}

	//var a float64 = 41
	//var b float64 = 30
	//a = (a + b) / 2.0
	//b = a*2 - b
	//a = a*2 - b
	//fmt.Println(a)
	//fmt.Println(b)

	var num = 11 //1011
	num1 := fmt.Sprintf("%b", num)
	fmt.Printf("%T", num1)

	num2 := num >> 1 //101 5
	num3 := num << 1 //10110 2+4+16=22

	fmt.Println(num2)
	fmt.Println(num3)

}
