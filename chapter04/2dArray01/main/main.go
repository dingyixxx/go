package main

import "fmt"

func main() {
	shi := [4][5]rune{}
	p := "床前明月光疑是地上霜举头望明月低头思故乡"
	arr := []rune(p)
	idx := 0
	for i := 0; i < len(shi); i++ {
		for j := 0; j < len(shi[0]); j++ {
			shi[i][j] = arr[idx]
			idx++
		}
	}

	//打印地址，用%p
	fmt.Printf("&(shi[0])的值是%p", &(shi[0]))

	//rune int32 4
	//0x1400011c000
	//0x1400011c010
	//0x1400011c014
	//0x1400011c028
	//0x1400011c03c
	fmt.Println(&(shi[0][0]))
	fmt.Println(&(shi[0][4]))
	fmt.Println(&(shi[1][0]))
	fmt.Println(&(shi[2][0]))
	fmt.Println(&(shi[3][0]))

	fmt.Println()

	//int64 8
	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(&arr2[0][0])
	fmt.Println(&arr2[1][0])
	//0x1400001a090
	//0x1400001a0a8

	//arr3 := [...][3]int{{1, 2, 3}, {4, 55, 66}}
	//arr3 := [2][...]int{{1, 2, 3}, {4, 55, 66}} 错误
	//arr3 := [...][...]int{{1, 2, 3}, {4, 55, 66}} 错误
	arr3 := [2][3]int{{1, 2, 3}, {4, 55, 66}}

	fmt.Println(arr3) //[[1 2 3] [4 55 66]]

}
