package main

import (
	"fmt"
	"runtime"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp    Point
	rightDown Point
}

type RectPtr struct {
	leftUpPtr    *Point
	rightDownPtr *Point //注意，此处和上面的Rect不一致，此处，都是指针
}

func main() {

	rect := Rect{Point{2, 2}, Point{10, 10}}
	//	结构体的所有字段的数据，在“内存”中是“连续”的

	fmt.Printf("rect-leftUp-x: %p,"+
		"rect-leftUp-y: %p,"+
		"rect-rightDown-x: %p,"+
		"rect-rightDown-y: %p\n",
		&rect.leftUp.x,
		&rect.leftUp.y, &rect.rightDown.x, &rect.rightDown.y)
	//该是连续
	//rect-leftUp-x: 0x140000180a0,rect-leftUp-y: 0x140000180a8,rect-rightDown-x: 0x140000180b0,rect-rightDown-y: 0x140000180b8

	fmt.Printf("rect-leftUp:%p, rect-rightDown:%p\n", &rect.leftUp, &rect.rightDown)
	//该是连续
	//rect-leftUp:0x140000180a0, rect-rightDown:0x140000180b0

	fmt.Println("------")
	MakeMemoryFragments()

	rect1 := RectPtr{&Point{2, 2}, &Point{10, 10}}

	//	rect1 RectPtr的属性值 是连续的，但是，它们指向的地址，不是连续的
	fmt.Printf("rect1-leftUpPtr本身的地址:%p, rect1-rightDownPtr本身的地址:%p\n",
		&rect1.leftUpPtr, &rect1.rightDownPtr)
	//该是连续
	//一个指针在 64 位机器上占 8 字节
	//rect1-leftUpPtr本身的地址:0x14000012020, rect1-rightDownPtr本身的地址:0x14000012028

	fmt.Printf("rect1-leftUpPtr指向的地址:%p, rect1-rightDownPtr指向的地址:%p\n",
		rect1.leftUpPtr, rect1.rightDownPtr)
	//未必连续，因为之前已经造了很多碎片
	//rect1-leftUpPtr指向的地址:0x14000010090, rect1-rightDownPtr指向的地址:0x140000100c0

	fmt.Printf("rect1-leftUpPtr指向的地址的x:%p, rect1-leftUpPtr指向的地址的y:%p, rect1-rightDownPtr指向的地址的x:%p, rect1-rightDownPtr指向的地址的y:%p\n",
		&((*rect1.leftUpPtr).x), &((*rect1.leftUpPtr).y), &((*rect1.rightDownPtr).x), &((*rect1.rightDownPtr).y))
	//前二者连续，后二者连续，而，前二者 vs 后二者，因为之前已经造了很多碎片
	//rect1-leftUpPtr指向的地址的x:0x14000010090, rect1-leftUpPtr指向的地址的y:0x14000010098, rect1-rightDownPtr指向的地址的x:0x140000100c0, rect1-rightDownPtr指向的地址的y:0x140000100c8

}

// 故意营造一种“碎片化”了的内存
func MakeMemoryFragments() {
	// 第1步：分配一堆小块，占满内存
	blocks := make([][]byte, 100)
	for i := 0; i < 100; i++ {
		blocks[i] = make([]byte, 9)
	}

	// 第2步：释放偶数位置的块，制造"碎片空洞"
	for i := 0; i < 100; i += 2 {
		blocks[i] = nil // 释放，产生碎片
	}
	runtime.GC() // 强制 GC

}

//	https://leetcode.com/problems/booking-concert-tickets-in-groups/description/
