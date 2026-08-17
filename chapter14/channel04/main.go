package main

import (
	"fmt"
	"time"
)

// 乒乓转发...利用了“nil channel会阻塞读/写”的特性...思路清奇...
//
//	但我应该不会这么写，因为感觉“交替打印FooBar”也可以实现这个需求，Foo-Bar-Foo-Bar-Foo-Bar
func main() {
	inch := make(chan int)
	outch := make(chan int)
	go func() {
		var in <-chan int = inch //只读，定义成常规channel的话，其实也是可以的
		var out chan<- int       //只写，定义成常规channel的话，其实也是可以的
		var val int
		for {
			select {
			case out <- val:
				out = nil
				in = inch
			case val = <-in:
				out = outch
				in = nil
			}
		}
	}()

	go func() {
		for r := range outch {
			fmt.Println("result:", r)
		}
	}()

	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3 * time.Second)
}
