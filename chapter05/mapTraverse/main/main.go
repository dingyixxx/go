package main

import "fmt"

func main() {
	//map不能用for的下标来遍历，只能通过for-range来遍历
	//形成条件反射，map上来就是make
	books := make(map[int]map[string]string)
	books[1] = make(map[string]string) //用=而不是:=
	books[1]["bookName"] = "西游记"
	books[1]["price"] = "13.88"
	books[1]["author"] = "吴承恩"

	books[2] = make(map[string]string)
	books[2]["bookName"] = "水浒传"
	books[2]["price"] = "25.67"
	books[2]["author"] = "施耐庵"

	//map的长度，也用len
	fmt.Println(len(books))
	fmt.Println(len(books[1]))

	for k1, _ := range books {
		fmt.Printf("当前遍历的是books[%v]\n", k1)
		for k2, v2 := range books[k1] {
			fmt.Printf("books[%v][%v]=%v\n", k1, k2, v2)
		}
	}
}
