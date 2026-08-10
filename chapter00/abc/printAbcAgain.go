package main

import (
	"fmt"
	"sync"
	"time"
)

var chanA chan int = make(chan int, 1)
var chanB chan int = make(chan int, 1)
var chanC chan int = make(chan int, 1)
var count int = 10
var wg sync.WaitGroup

func A() {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanA
		fmt.Print("A")
		time.Sleep(time.Second)
		chanB <- 1

	}
}
func B() {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanB
		fmt.Print("B")
		time.Sleep(time.Second)
		chanC <- 1

	}
}
func C() {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanC
		fmt.Print("C")
		time.Sleep(time.Second)
		fmt.Println()
		chanA <- 1

	}
}

func main() {
	wg.Add(3)
	go A()
	go B()
	go C()
	chanA <- 1
	wg.Wait()
	fmt.Println("都打印完了")

}
