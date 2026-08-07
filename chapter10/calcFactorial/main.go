package main

import (
	"fmt"
	"sync"
)

var m = make(map[uint64]uint64)
var lock sync.Mutex

func main() {
	var i uint64 = 1
	for i = 1; i <= 400; i++ {
		go calcFactorial(i)
	}
	lock.Lock() //todo 溢出 加锁
	for k, v := range m {
		fmt.Printf("%v的阶乘是:%v\n", k, v)
	}
	lock.Unlock()
}

func calcFactorial(n uint64) {
	var res uint64 = 1
	var i uint64 = 1
	for i = 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	m[n] = res
	lock.Unlock()
}
