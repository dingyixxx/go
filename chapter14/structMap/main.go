package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]*data{"x": {"one"}}
	m["x"].name = "two" //error
	fmt.Println(*(m["x"]))

	//m := map[string]data{"x": {"one"}}
	//m["x"].name = "two" //error Cannot assign to m["x"].name
}
