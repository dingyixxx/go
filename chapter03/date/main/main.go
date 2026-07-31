package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Unix())     //unix时间戳
	fmt.Println(now.UnixNano()) //纳秒时间戳

	//	获取随机数

}
