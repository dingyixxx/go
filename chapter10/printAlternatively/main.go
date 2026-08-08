package main

import (
	"fmt"
	"time"
)

var aChan chan int = make(chan int, 1)
var bChan chan int = make(chan int, 1)
var cChan chan int = make(chan int, 1)
var count int = 60
var doneChan chan int = make(chan int, count*999)
var doneCount chan int = make(chan int, count*3)

func printA() {
	for {
		_, ok := <-aChan
		if ok == true {
			fmt.Print("a")
			bChan <- 1
			doneChan <- 1
			doneCount <- 1
			time.Sleep(time.Millisecond * 80)
		} else {
			fmt.Println("printA结束...")
			break
		}
	}
}

func printB() {
	for {
		_, ok := <-bChan
		if ok == true {
			fmt.Print("b")
			cChan <- 1
			doneChan <- 1
			doneCount <- 1

			time.Sleep(time.Millisecond * 99)
		} else {
			fmt.Println("printB结束...")
			break
		}

	}
}

func printC() {
	for {
		_, ok := <-cChan
		if ok == true {
			fmt.Print("c")
			fmt.Println()
			doneChan <- 1
			doneCount <- 1
			if len(doneCount) != count*3 {
				aChan <- 1
			}

			//fmt.Printf("len(doneCount)-%v\n", len(doneCount))
			//
			//fmt.Printf("len(doneChan)-%v\n", len(doneChan))
			//每次打印出来,都是0,
			//因为main一直在实时消费doneChan,
			//即主协程"看到有值就拿走,不会等着doneChan全写完"

			time.Sleep(time.Millisecond * 100)
		} else {
			fmt.Println("printC结束...")
			break
		}
	}
}

func main() {
	aChan <- 1
	go printC()
	go printB()
	go printA()

	total := count * 3
	for i := 0; i < total; i++ {
		<-doneChan
	}
	close(cChan)
	close(bChan)
	close(aChan)

	time.Sleep(time.Millisecond * 900)

	fmt.Println("end...")
}
