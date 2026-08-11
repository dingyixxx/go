package main

import (
	"fmt"
	"sync"
	"time"
)

//waitGroup=countDownLatch

func worker(id int, wg *sync.WaitGroup) { //注意点四：传入引用
	defer wg.Done() //注意点二：子协程任务完成之后，报告完成
	fmt.Printf("worker-%v开始工作...\n", id)
	time.Sleep(time.Second * 4)
	fmt.Printf("worker-%v结束工作...\n", id)
}
func main() {
	count := 6
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1) //注意点一：不能在子协程里面再add，先add，才能够让主协程阻塞住
		go worker(i, &wg)
	}
	wg.Wait() //注意点三：必须等到所有的都完成，才算完成
	fmt.Println("全部都做完了...")
}
