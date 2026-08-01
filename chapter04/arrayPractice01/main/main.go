package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	letters := [26]byte{}
	for i := 0; i < 26; i++ {
		letters[i] = byte('A' + i)
		//letters[i] = byte(65 + i)

	}

	for _, l := range letters {
		//letter := fmt.Sprintf("%c", l)
		//fmt.Println(letter)
		fmt.Printf("%c", l)
	}
	fmt.Println()

	intArr := [5]int{1, -1, 9, 90, 12}
	intArrD := decimal.NewFromInt(int64(intArr[0]))
	for i := 1; i < len(intArr); i++ {
		intArrD = intArrD.Add(decimal.NewFromInt(int64(intArr[i])))
	}
	intArrD = intArrD.DivRound(decimal.NewFromInt(int64(len(intArr))), 4)
	//f, exact := intArrD.Float64()
	//只要结果的小数部分不是 0.5、0.25、0.125 等"2的幂分之一"，
	//Float64() 的 exact 就一定是 false。
	//因此，无需打印
	fmt.Println(intArrD)
}
