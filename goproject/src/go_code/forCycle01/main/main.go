package main

import "fmt"

func main() {
	//	打印一个菱形

	layers := 5
	for i := 0; i < layers; i++ {
		for j := 0; j < layers; j++ {
			if i == layers-1-j {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}

		for j := 0; j < layers-1; j++ {
			if i == j+1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < layers-1; i++ {
		for j := 0; j < layers; j++ {
			if i == j-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		for j := 0; j < layers-1; j++ {
			if i == layers-3-j {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
