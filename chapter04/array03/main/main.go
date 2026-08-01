package main

import "fmt"

func main() {
	//	numArrNew接收的地址0x140001a6018
	//	test01接收的地址0x140001a6048
	//  [888 999 777]
	//	testPtr接收的地址0x140001a6018 可以看得出来 指针接收的地址就是入参的地址
	//	[909090 909090 909090]
	var numArrNew = [3]int{888, 999, 777}
	fmt.Printf("numArrNew接收的地址%p\n", &numArrNew)
	//打印内存地址的占位符是%p

	test01(numArrNew)
	fmt.Println(numArrNew)

	testPtr(&numArrNew) //传指针
	fmt.Println(numArrNew)
	//	5种初始化数组的方式
	var numArr1 = [3]int{1, 2, 3}
	var numArr2 = [4]int{4, 5, 6, 7}
	var numArr3 = [...]int{11, 12, 13, 14, 15}
	var numArr4 = [...]int{2: 2222, 1: 1111}
	//类型推导
	numArr5 := [...]string{1: "do", 4: "fa", 6: "ra", 7: "ti"}

	fmt.Println(numArr1)
	fmt.Println(numArr2)
	fmt.Println(numArr3)
	fmt.Println(numArr4)

	// 快捷键forr for range
	for i, i2 := range numArr5 {
		fmt.Printf("%d-%q  \n", i, i2)
		fmt.Printf("numArr5[%d]=%v  \n", i, numArr5[i])
	}

	//	go的int里只能放int，不像python什么都能放
	//	数组长度不能变
	//strArr :=[6]string{}
	//strArr[0]=13.1
	//strArr[6]="hello"

	//int64的数组，不能放int32
	//numArr :=[...]int32{11,12}
	//numArr[0]=int64(2)

	//[]必须要写大小，否则定义的就是切片
	//strArr :=[6]string{}

	//	非常重要的一点是：数组是值拷贝，非常重要，因此，务必要看是否需要用指针传递
	//《飞云之上》韩红

}

// 定义函数时，参数类型，[3]int 和 [4]int 是两种不一样的数据类型，务必要注意，也会编译不过去
// 如果不写3,即如果写成了[]int，那就是切片
// 不能把数组传给接收“切片”作为入参的函数
func test01(arr [3]int) {
	fmt.Printf("test01接收的地址%p\n", &arr)
	arr[0] = 909090
	arr[1] = 909090
	arr[2] = 909090
}

// 接收指针
func testPtr(arr *[3]int) {
	fmt.Printf("testPtr接收的地址%p\n", arr)
	arr[0] = 909090
	arr[1] = 909090
	arr[2] = 909090
}
