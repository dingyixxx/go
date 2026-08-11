package main

import (
	"fmt"
	"sync"
	"time"
)

var chanA = make(chan int, 1)
var chanB = make(chan int, 1)
var chanC = make(chan int, 1)
var wg sync.WaitGroup
var count int = 6

func sayA() {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanA
		time.Sleep(time.Second * 1)
		fmt.Print("a")
		chanB <- 1
	}
}
func sayB() {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanB
		time.Sleep(time.Second * 1)
		fmt.Print("b")
		chanC <- 1
	}
}
func sayC() {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanC
		time.Sleep(time.Second * 1)
		fmt.Print("c")
		fmt.Println()
		chanA <- 1
	}
}
func main() {
	wg.Add(3)
	chanA <- 1
	go sayA()
	go sayB()
	go sayC()
	wg.Wait()
	fmt.Printf("打印完了%v轮abc...", count)
}
