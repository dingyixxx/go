package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close() //关闭redis连接
	_, err = c.Do("Set", "key1", "hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.String(c.Do("Get", "key1"))
	if err != nil {
		fmt.Println("get key1 failed,", err)
		return
	}
	fmt.Println(r)

}
