package main

import "fmt"

func main() {
	floatArr := [6]float64{}
	for i := 0; i < len(floatArr); i++ {
		fmt.Scanf("%f", &floatArr[i])
	}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("数组floatArr的第%d项是%.4f\n", i+1, floatArr[i])
	}

}
