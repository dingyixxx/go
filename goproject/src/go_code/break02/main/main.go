package main

import "fmt"

func main() {

label_outer:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label_outer //break可以手工指定跳出哪个标签对应的循环
			}
			fmt.Println("j=", j, "i=", i)
		}
	}

	//break默认跳出离它最近的for循环
	//for i := 0; i < 4; i++ {
	//	for j := 0; j < 10; j++ {
	//		if j == 2 {
	//			break
	//		}
	//		fmt.Println("j=", j, "i=", i)
	//	}
	//}
}
