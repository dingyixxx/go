package main

import "fmt"

var done bool = false

func main() {
	go func() {
		done = true
	}()
	for !done {
	}
	fmt.Println("done!") //GODEBUG=asyncpreemptoff=1;GOMAXPROCS=1
}
