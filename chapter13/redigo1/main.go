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
	_, err = c.Do("HSet", "user01", "name", "汤姆")

	//这里需要根据name对应的类型来使用redis.Xxx的方法
	//如果存放的是int则应当使用 redis.Int()，看相关手册
	r, err := redis.String(c.Do("HGet", "user01", "name"))

	_, err = c.Do("MSet", "name", "you123", "address", "北京昌平~")
	r1, err := redis.Strings(c.Do("MGet", "name", "address"))

	for _, v := range r1 {
		fmt.Println(v)
	}

	fmt.Println(r)

}
