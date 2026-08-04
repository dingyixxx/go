package main

import "fmt"

type Person struct {
	Name string
}

func (self Person) speak() {
	fmt.Printf("%v是个神\n", self.Name)
}
func (self Person) jisuan() {
	num := 0
	for i := 1; i <= 1000; i++ {
		num += i
	}
	fmt.Println("从1加到1000", num)
}
func (self Person) jisuan2(n int) {
	num := 0
	for i := 1; i <= n; i++ {
		num += i
	}
	fmt.Printf("%v加到%v=%v\n", 1, n, num)
}
func (self Person) getSum(a int, b int) (sum int) {
	fmt.Printf("%v+%v=%v\n", a, b, a+b)
	return (a + b)
}

func main() {
	p := Person{
		Name: "三和",
	}
	p.speak()
	p.jisuan()
	p.jisuan2(1000)
	fmt.Println(p.getSum(33, 45))
}
