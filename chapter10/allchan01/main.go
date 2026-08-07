package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "tom~", Age: 180}
	allChan <- 10
	allChan <- "jack"
	allChan <- cat1
	allChan <- cat2

	<-allChan
	<-allChan
	// 取出 — 需要类型断言才能访问具体类型的字段
	cat11 := <-allChan
	cat12 := cat11.(Cat)
	fmt.Println(cat12.Age)
	c, ok := cat11.(Cat) // 类型断言：将 interface{} 转为 Cat
	if ok {
		fmt.Println(c.Name) // ✅ 输出: tom
	} else {
		fmt.Println("不是 Cat 类型")
	}
}
