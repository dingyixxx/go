package main

import "fmt"

func main() {
	//	打印九九乘法表
	layers := 9
	for i := 1; i <= layers; i++ {
		for j := 1; j <= layers; j++ {
			if j <= i {
				fmt.Printf("%v * %v = %v \t", j, i, i*j)
			}
		}
		fmt.Println("")
	}
}
