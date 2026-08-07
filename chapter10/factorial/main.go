package main

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

var count int64 = 200

// 如果是200个协程，如果不加锁，那是并发写，就会报这个错误
var myMap map[int64]decimal.Decimal = make(map[int64]decimal.Decimal, count)

var lock = sync.Mutex{}

// 只要涉及到阶乘，就不要想用int或者uint了，就直接用decimal，除非是20及其以下的阶乘
func main() {
	for i := int64(1); i <= count; i++ {
		go GenFactorial(i)
	}
	//问题二：等多久，先拍个睡眠时长
	//time.Sleep(time.Second * 60)

	//问题一：并发写 解法：加锁
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("myMap的key %v 的 value 是： %v\n", k, v)
	}

	//把打印map键值对的方法前后的加解锁注释掉，
	//然后执行go run -race main.go这个命令，
	//会出现Found 6 data race(s)

	//把写入map的加解锁注释掉，
	//然后
	//第一种操作：执行go run -race main.go这个命令，
	//会出现Found 7 data race(s)
	//第二种操作：直接执行
	// fatal error: concurrent map writes

	//但是，加完锁，也不会“等到”200个协程都执行完毕再打印出来，只能打印出来40-80多个key
	//暂且先sleep一下,60秒基本能稳定地写完200条记录
	lock.Unlock()
}

func GenFactorial(num int64) (res decimal.Decimal) {
	res = decimal.NewFromInt(1)
	for i := int64(1); i <= num; i++ {
		res = res.Mul(decimal.NewFromInt(i))
	}
	lock.Lock()
	myMap[num] = res
	lock.Unlock()
	return
}
