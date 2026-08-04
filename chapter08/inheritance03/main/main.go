package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	tv := TV{
		Goods{"电视机", 1111.23},
		Brand{"海尔", "山东"},
	}
	fmt.Println(tv)
	fmt.Println(tv.Brand.Name)
	fmt.Println()
	tv2 := &TV2{
		&Goods{"冰箱", 88.679},
		&Brand{"格兰仕", "上海"},
	}
	fmt.Println(tv2)
	fmt.Println((*(*tv2).Brand).Name)
}
