package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		//time.Sleep(time.Second *0.1)//Invalid operation: time.Second *0.1 (cannot convert the constant 0.1 to the type int64)
		time.Sleep(100 * time.Millisecond)
		if i == 100 {
			break
		}
	}
}
