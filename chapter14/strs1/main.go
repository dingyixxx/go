package main

import "fmt"

func main() {
	done := false
	go func() {
		done = true
	}()
	for !done {
		fmt.Println("not done!")
		fmt.Println("not done!")
		fmt.Println("not done!")
		fmt.Println("not done!")
		fmt.Println("not done!")
	}
	fmt.Println("done!") //GODEBUG="asyncpreemptoff=1";GOMAXPROCS=1
}
