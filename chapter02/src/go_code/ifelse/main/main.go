package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入成绩")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("奖励100分选手")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台(80，99]的iphone7plus")
	} else if score > 60 && score <= 80 {
		fmt.Println("奖励一个(60，80]的ipad")
	} else {
		fmt.Println("没有什么奖励")
	}
}
