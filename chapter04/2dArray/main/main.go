package main

import "fmt"

func main() {
	arr := [10][20]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if i%5 == j%5 {
				arr[i][j] = 9999
			} else {
				arr[i][j] = 1
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%d\t\t", arr[i][j])
		}
		fmt.Println()
	}

}
