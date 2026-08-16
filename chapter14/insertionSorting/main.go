package main

import "fmt"

func main() {
	//插入排序

	//insertion sorting的基本思想是：
	//把n个待排序的元素，看成为一个有序表和一个无序表，

	//开始时，
	//有序表中只包含一个元素，无序表中包含有n-1个元素，

	//排序过程中，
	//每次从无序表中取出第一个元素，把它的排序码依次与有序表元素的排序码进行比较，
	//注意，是从有序数组的“后”，开始比较
	//将它插入到有序表中的适当位置，使之成为新的有序表

	arr := []int{20, 1, 11, 102, 33, -9, -13, 77, -66}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		valToBeInserted := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > valToBeInserted {
			arr[j+1] = arr[j] //给valToBeInserted腾挪出一个位置
			j--
		}
		arr[j+1] = valToBeInserted
	}
}
