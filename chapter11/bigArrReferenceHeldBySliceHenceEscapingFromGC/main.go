package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//Arrays.asList()
	// 1. 创建一个 100MB 的大数组
	bigArray := make([]byte, 100*1024*1024)
	for i := range bigArray {
		bigArray[i] = 'A'
	}
	fmt.Printf("大数组创建完成, 大小: %d MB\n", len(bigArray)/1024/1024)

	// 2. 只取前 10 个字节
	smallSlice := bigArray[:10]
	fmt.Printf("smallSlice 长度: %d, 但底层数组: %d MB\n",
		len(smallSlice), cap(bigArray)/1024/1024)

	// 3. 把大数组变量置 nil（以为能回收）
	bigArray = nil

	// 4. 强制 GC
	runtime.GC()
	printMem()

	// 5. 但 smallSlice 还活着，底层 100MB 数组无法回收
	fmt.Println("smallSlice 还活着，100MB 无法回收...")
	time.Sleep(time.Second)

	// 6. 验证：smallSlice 仍能访问（虽然只用了10字节）
	fmt.Printf("smallSlice[0] = %c\n", smallSlice[0])

	// 7. 只有 smallSlice 也失效后，内存才能回收
	smallSlice = nil
	runtime.GC()
	printMem()
	fmt.Println("现在 100MB 才被回收")
}

func printMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("当前内存: %d MB\n", m.Alloc/1024/1024)
}
