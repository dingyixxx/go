package main

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
}

func main() {
	var c C
	//c.Name
	c.A.Name = "jacky"

}
