package main

import "fmt"

func main() {
	arr := []int{24, 69, 80, 57, 13}
	bubbleSort(arr)
	fmt.Println("arr-", arr)

	// cmd+alt+m 抽象为函数
	arr1 := [7]int{24, 69, 80, 11, 57, 13, 7}
	bubbleSortArray(&arr1)
	fmt.Println("arr1-", arr1)
}

func bubbleSortArray(arrPtr *[7]int) {
	len := len(*arrPtr)
	for i := 0; i < len; i++ {
		for j := 0; j < (len - i - 1); j++ {
			//给指针整体，加一个括号
			if (*arrPtr)[j] < (*arrPtr)[j+1] {

			} else {
				swapArray(arrPtr, j, j+1)
			}
		}
	}
}

func swapArray(arrPtr *[7]int, j int, i int) {
	temp := (*arrPtr)[i]
	(*arrPtr)[i] = (*arrPtr)[j]
	(*arrPtr)[j] = temp
}

// 切片是引用类型，数组是值类型，所以，切片可以直接传，数组必须传指针
// (用数组的话，
// 一个是函数要传入“定长”的数组，
// 二是函数内部、不写数组引用的话、编译器也不会报错、因此实在令人不能确定：是否真真实实地改变了“原数组”）
func bubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := 0; j < (len - i - 1); j++ {
			if arr[j] < arr[j+1] {

			} else {
				swap(arr, j, j+1)
			}
		}
	}
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
