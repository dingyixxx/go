package main

import (
	"fmt"
)

func main() {
	//	如果函数的入参是“数组的引用”，则，
	//	函数内部，凡是涉及到解指针，就要用*，比如，排序

	//golang的函数，首字母，尽量大写

	arr := [9]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	var idx int = BinarySearch(&arr, 80)
	fmt.Println(idx)

}

func BinarySearch(arr *[9]int, target int) int {

	//这里传给BinarySearchWithin的入参，不应该是*arr
	//只有解指针的时候，再用*
	return BinarySearchWithin(arr, 0, len(arr)-1, target)
}

func BinarySearchWithin(arr *[9]int, from int, to int, target int) int {
	if arr[to] < target || arr[from] > target {
		return -1
	}
	mid := (from + to) / 2
	if (*arr)[mid] < target {
		return BinarySearchWithin(arr, mid+1, to, target)
	} else if (*arr)[mid] > target {
		return BinarySearchWithin(arr, from, mid, target)
	} else {
		return mid
	}
}
