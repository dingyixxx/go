package main

import "fmt"

func main() {
	//切片是引用类型，在进行传递时，遵守引用传递的机制

	//切片是一个可以动态变化的数组
	var slice1 = []int{1, 2, 3}
	slice1 = append(slice1, 4)
	fmt.Println(slice1)

	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := intArr[1:3]
	fmt.Println(slice2)

	fmt.Println("intArr=", intArr)
	fmt.Println("slice2=", slice2)
	fmt.Println("slice2的元素个数=", len(slice2))
	//数组长度-起始下标
	//继续放多少个，才会触发“底层数组”扩容
	fmt.Println("slice2的元素容量=", cap(slice2))

	slice2[0] = 22222222222
	fmt.Println("intArr=", intArr)
	fmt.Println("slice2=", slice2)
	//intArr= [1 22222222222 3 4 5 6 7 8 9]
	//slice2= [22222222222 3]
	intArr[2] = 333333333333
	fmt.Println("intArr=", intArr)
	fmt.Println("slice2=", slice2)
	//intArr = [1 22222222222 333333333333 4 5 6 7 8 9]
	//slice2 = [22222222222 333333333333]

	//slice2[0]-0x1400009e008
	//intArr[1]-0x1400009e008
	fmt.Printf("slice2[0]-%v\n", &slice2[0])
	fmt.Printf("intArr[1]-%v\n", &intArr[1])

	//slice从底层来说，其实就是一个数据结构（struct结构体）
	//type slice struct{
	//	ptr *[2]int
	//	len int
	//	cap int
	//}

}
