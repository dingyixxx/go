package main

import "fmt"

func main() {

	//好不容易在javascript里面记死了不允许使用的var，又在学golang的时候给强制用上了

	//	continue也可以指定，跳出哪一轮循环

label_outer:
	for i := 0; i < 6; i++ {
		fmt.Printf("外层循环的值i是：%v ------\n", i)
		for j := 0; j < 4; j++ {
			if j == 1 {
				//continue
				continue label_outer
			}
			fmt.Printf("内层循环j的值是：%v \n", j)
		}
	}

	//	0 -0 1 2
}
