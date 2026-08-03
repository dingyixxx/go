package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	chA := make(chan struct{})
	chB := make(chan struct{})
	chC := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			<-chA
			time.Sleep(time.Second)
			fmt.Print("a")
			chB <- struct{}{} // A 每轮都发
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			<-chB
			time.Sleep(time.Second)
			fmt.Print("b")
			chC <- struct{}{} // B 每轮都发
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			<-chC
			time.Sleep(time.Second)
			fmt.Print("c")
			if i < 2 { // 只有 C 最后一轮不发
				chA <- struct{}{}
			}
		}
	}()

	chA <- struct{}{}
	fmt.Println("...")
	wg.Wait()
	fmt.Println("打印完了...")
}
