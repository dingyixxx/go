package main

import "fmt"

var slice = []int{1, 1}

func fib(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		if num < len(slice) && slice[num] != 0 {
			return slice[num]
		}
		var res = fib(num-1) + fib(num-2)
		slice = append(slice, res)
		return res
	}
}

func main() {
	fib(9)
	fmt.Println(slice)
}
