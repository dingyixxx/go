package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 0
	num := 0
	for {
		rand.Seed(time.Now().Unix())
		num = rand.Intn(100) + 1
		fmt.Printf("本次生成的随机数是：%d\n", num)
		count++
		if num == 99 {
			break
		}
	}

	fmt.Printf("最终生成了%d次随机数\n", count)
}
