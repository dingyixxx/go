package main

import (
	"fmt"
)

func main() {
	done := false
	go func() {
		done = true
	}()
	for !done {
		//runtime.Gosched()  //GODEBUG="asyncpreemptoff=1";GOMAXPROCS=1
	}
	fmt.Println("done!")
}
