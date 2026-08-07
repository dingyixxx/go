package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// 1) chan Cat（值类型）
	//var catChan chan Cat
	//catChan = make(chan Cat, 10)
	//
	//cat1 := Cat{Name: "tom", Age: 18}
	//cat2 := Cat{Name: "tom-", Age: 180}
	//catChan <- cat1
	//catChan <- cat2
	//
	//cat11 := <-catChan
	//cat22 := <-catChan
	//
	//fmt.Println(cat11, cat22)
	//
	//// 2) chan *Cat（指针类型）
	//var catChan2 chan *Cat
	//catChan2 = make(chan *Cat, 10)
	//
	//cat3 := Cat{Name: "tom", Age: 18}
	//cat4 := Cat{Name: "tom~", Age: 180}
	//catChan2 <- &cat3
	//catChan2 <- &cat4
	//
	//cat33 := <-catChan2
	//cat44 := <-catChan2
	//
	//fmt.Println(cat33, cat44)

	// 3) chan interface{}（任意数据类型）
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	cat5 := Cat{Name: "tom", Age: 18}
	cat6 := Cat{Name: "tom~", Age: 180}
	allChan <- cat5
	allChan <- cat6
	allChan <- 10
	allChan <- "jack"

	// 取出
	cat111 := <-allChan
	cat222 := <-allChan
	v1 := <-allChan
	v2 := <-allChan

	fmt.Println(cat111, cat222, v1, v2)
}
