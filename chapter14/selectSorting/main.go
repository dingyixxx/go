package main

import "fmt"

func main() {
	//	选择排序
	//	start元素和所有其后元素全局比较、小的留下、范围的start逐渐变大
	arr := []int{20, 1, 11, 102, 33, -9, -13, 77, -66}
	selectSorting(arr)
	fmt.Println(arr)

	//	冒泡排序
	//	和相邻元素比较、大泡冒上去、范围的end逐渐变小
	arr1 := []int{20, 1, 11, 102, 33, -9, -13, 77, -66}
	bubbleSort(arr1)
	fmt.Println(arr1)
}

func selectSorting(arr []int) {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		minIdx := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

}

func bubbleSort(arr []int) {
	lengh := len(arr)
	for i := 0; i < lengh; i++ {
		for j := 0; j < lengh-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
