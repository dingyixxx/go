package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// ❌ 错误：slice 持有大数组引用
func getBad() []byte {
	raw := make([]byte, 10000)
	raw[0] = 'A'
	raw[1] = 'B'
	raw[2] = 'C'
	return raw[:3] // 引用 raw 的底层数组
}

// ✅ 正确：copy 到新数组
func getGood() []byte {
	raw := make([]byte, 10000)
	raw[0] = 'A'
	raw[1] = 'B'
	raw[2] = 'C'
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}

func printMem(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Alloc=%d KB, Sys=%d KB\n", label, m.Alloc/1024, m.Sys/1024)
}

func main() {
	debug.SetGCPercent(-1) // 禁用自动 GC，手动控制

	fmt.Println("=== 错误写法：slice 引用大数组 ===")
	printMem("初始")
	var refs [][]byte
	for i := 0; i < 1000; i++ {
		refs = append(refs, getBad())
	}
	printMem("创建1000个slice后")
	refs = nil // 释放引用
	runtime.GC()
	printMem("GC后（大数组才被回收）")

	fmt.Println("\n=== 正确写法：copy 到新数组 ===")
	printMem("初始")
	var copies [][]byte
	for i := 0; i < 1000; i++ {
		copies = append(copies, getGood())
	}
	printMem("创建1000个slice后")
	copies = nil // 释放引用
	runtime.GC()
	printMem("GC后（大数组已回收）")
}
