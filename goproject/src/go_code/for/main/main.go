package main

import "fmt"

func main() {
	//i := 0
	//for ; i < 10; i++ {
	//	fmt.Printf("你好,我是第%d个小苹果\n", i+1)
	//}

	//for i < 10 {
	//	fmt.Printf("你好,我是第%d个小苹果\n", i+1)
	//	i++ //写这里也可以
	//}
	//fmt.Println(i)

	//start := time.Now()
	//for {
	//	fmt.Println("hello")
	//	if time.Since(start) >= 5*time.Second {
	//		break
	//	}
	//}
	//fmt.Println("5秒已到，跳出循环！")

	//i := 0
	//for {
	//	if i < 10 {
	//		fmt.Println("ok~")
	//	} else {
	//		break
	//	}
	//	i++
	//}
	//fmt.Println("跳出循环~")

	i := 0
	for {
		//熟悉的;;
		if i < 10 {
			fmt.Println("ok~")
		} else {
			break
		}
		i++
	}
	fmt.Println("跳出循环~")

}
