package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	var balance float64 = 100000
	var count = 0
	var balanceD = decimal.NewFromFloat(balance)
	for {
		if balanceD.GreaterThan(decimal.NewFromFloat(50000)) {
			balanceD = balanceD.Mul(decimal.NewFromFloat(0.95))
		} else {
			balanceD = balanceD.Sub(decimal.NewFromInt(1000))
		}
		count++
		if balanceD.LessThan(decimal.Zero) {
			break
		}
	}
	fmt.Println(count)

}
