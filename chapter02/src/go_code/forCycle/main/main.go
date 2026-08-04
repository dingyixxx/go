package main

import "fmt"

func main() {
	layers := 9
	for i := 0; i < layers; i++ {
		for j := 0; j < layers+1; j++ {
			if j == layers-i || (i == layers-1 && j >= layers-i && j%2 == 1) {
				fmt.Printf("❄")
			} else {
				fmt.Printf(" ")
			}
		}

		for j := 0; j < layers; j++ {
			if j == i-1 || (i == layers-1 && j <= i-1 && j%2 == 1) {
				fmt.Printf("❄")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

	//.........*.........
	//........*.*........
	//.......*...*.......
	//......*.....*......
	//.....*.......*.....
	//....*.........*....
	//...*...........*...
	//..*.............*..
	//.*...............*.

}
