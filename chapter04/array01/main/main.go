package main

import "fmt"

func main() {
	var arr [6]int
	//	默认值为0

	fmt.Println(arr)
	fmt.Println(&arr) //这么打印数组地址，是打印不出来的？
	fmt.Println(&arr[0])
	//数组的地址，和数组第一个元素的地址，是一样的
	fmt.Printf("arr的地址%p，arr[0]的地址%p\n", &arr, &arr[0])

	//arr的地址0x140000b4000，arr[0]的地址0x140000b4000，arr[1]的地址0x140000b4008 占8个字节
	fmt.Printf("arr的地址%p，arr[0]的地址%p，arr[1]的地址%p\n", &arr, &arr[0], &arr[1])
	fmt.Println()

	var arr1 [6]int32
	//arr1的地址0x140000a8018，arr1[0]的地址0x140000a8018，arr1[1]的地址0x140000a801c，arr1[2]的地址0x140000a8020，arr1[3]的地址0x140000a8024
	fmt.Printf("arr1的地址%p，arr1[0]的地址%p，arr1[1]的地址%p，arr1[2]的地址%p，arr1[3]的地址%p\n", &arr1, &arr1[0], &arr1[1], &arr1[2], &arr1[3])

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
