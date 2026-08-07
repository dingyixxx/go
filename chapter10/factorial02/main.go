package main

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

var count int64 = 200

var myMap map[int64]decimal.Decimal = make(map[int64]decimal.Decimal, count)

var intChan chan int64 = make(chan int64, count)
var exitChan chan bool = make(chan bool, 1)
var lock = sync.Mutex{}

func main() {
	for i := int64(1); i <= count; i++ {
		go GenFactorial(i)
	}

	for {
		_, ok := <-exitChan
		if ok == false {
			lock.Lock()
			for k, v := range myMap {
				//goroutine对于资源的“读”也有竞态检测，不像Java，随便读
				fmt.Printf("myMap的key %v 的 value 是： %v\n", k, v)
			}
			lock.Unlock()
			break
		}
	}

}

func GenFactorial(num int64) (res decimal.Decimal) {
	res = decimal.NewFromInt(1)
	for i := int64(1); i <= num; i++ {
		res = res.Mul(decimal.NewFromInt(i))
	}
	lock.Lock()
	myMap[num] = res
	intChan <- num
	fmt.Println("len(intChan)-", len(intChan))
	//记死了，加两个chan
	if len(intChan) == int(count) {
		close(intChan)
		exitChan <- true
		close(exitChan)
	}
	lock.Unlock()
	return
}
