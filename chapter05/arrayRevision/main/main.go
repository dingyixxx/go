package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	test01(&arr1)
	fmt.Println(arr1)
}

func test01(arr1 *[5]int) {
	(*arr1)[0] = 999
	(*arr1)[1] = 999
	(*arr1)[2] = 999
	(*arr1)[3] = 999
	(*arr1)[4] = 999
}
