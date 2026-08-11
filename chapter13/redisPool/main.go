package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:     8, // 最大空闲连接数
		MaxActive:   0, // 表示和数据库的最大链接数
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	//先从pool 取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "汤姆猫~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("r=", r)
}
