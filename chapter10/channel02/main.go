package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 10000)

	// chan 的make第二个参数传入什么数字，打印出来的cap(intChan)就是什么数字
	// 和 slice 不同，管道的容量在创建时就固定了，不会自动增长。
	fmt.Printf("intChan len(intChan) 是 %v cap(intChan) 是 %v\n", len(intChan), cap(intChan))

	//intChan len(intChan) 是 0 cap(intChan) 是 10000
	for i := 0; i < cap(intChan); i++ {
		//fmt.Printf("i=%v \n", i)
		intChan <- i + 1
	}
	for i := 0; i < 1000; i++ {
		<-intChan
	}
	fmt.Printf("intChan len(intChan) 是 %v cap(intChan) 是 %v\n", len(intChan), cap(intChan))
	//intChan len(intChan) 是 9000 cap(intChan) 是 10000
	//10000个元素塞进去了，取出来1000个元素，此时，只剩下9000个元素了，但是capacity还是10000个元素

	//为什么不能用cap(intChan)来遍历
	//cap(intChan)是10000，但此时，元素只有9000个，遍历到最后，就会报错
	//fatal error: all goroutines are asleep - deadlock!
	//因为，没有元素了，拿不出来了
	//for i := 0; i < cap(intChan); i++ {
	//	num := <-intChan
	//	fmt.Println("取出来的值是：", num)
	//}

	//为什么不能用len(intChan)来遍历
	//因为len(intChan)是动态的，每次从管道拿出来一个，len也少了一个
	//[a b c d e f g h] len为8
	//先拿出来a，len变为7，i变为1
	//再拿出来b，len变为6，i变为2
	//再拿出来c，len变为5，i变为3
	//再拿出来d，len变为4，i变为4
	//所以，只能拿出来当前队列里面，前一半的元素
	//就像Java的ArrayList里面，如果要删除元素，从0开始遍历，是一个典型的错
	//for i := 0; i < len(intChan); i++ {
	//	num := <-intChan
	//	fmt.Println("取出来的值是：", num)
	//}

	//不建议用普通的for循环（除非循环体内部不增加元素，且，先把len(intChan)在一开始的长度，取出来了）
	//lenth := len(intChan)
	//for i := 0; i < lenth; i++ {
	//	num := <-intChan
	//	fmt.Println("取出来的值是：", num)
	//}

	//因此，遍历管道，最好是用for range
	//管道没有下标，不能取出队列里面的第三个值
	//for v := range intChan {
	//	fmt.Println(v)
	//}
	//如果，在遍历之前，没有把管道给关了，取到最后，就会死锁
	//fatal error: all goroutines are asleep - deadlock!

	//死锁，是个报错，不是常说的“死锁”

	//因此，必须，在遍历之前，就把管道给关了，以实现“遍历完后，就退出遍历”
	close(intChan)
	for v := range intChan {
		fmt.Println(v)
	}
}
