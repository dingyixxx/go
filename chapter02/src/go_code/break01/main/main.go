package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机生成1-100的整数
	//返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic
	//go的rand包生成的随机数，也是不包括上界的，这一点和Java一样，和python不一样

	//记住，是math包里面的rand方法，不是别的包的rand方法

	//我们为了生成一个随机数，还需要给rand设置一个种子

	//所以那些"每次输出一样的值"的教程，都是基于 Go 1.20 之前的版本写的，你现在不用管这个问题

	//rand.Seed(time.Now().Unix())
	num := rand.Intn(100) + 1
	fmt.Println(num)

	fmt.Println(time.Now().Unix()) //time.Now().Unix()是时间戳，每次seed都变化，所以输出的随机数会变

}
