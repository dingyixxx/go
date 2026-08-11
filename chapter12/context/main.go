package main

import (
	"context"
	"fmt"
	"time"
)

//context貌似是个ThreadLocal+completableFuture的timeout的概念，
//也类似于js的cancelToken/abortController

func main() {
	// 1. 创建带超时的 context（3秒）
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 2. 在 context 上绑定值
	ctx = context.WithValue(ctx, "userId", 1001)
	ctx = context.WithValue(ctx, "role", "admin")

	go worker(ctx, "协程1")
	go worker(ctx, "协程2")

	time.Sleep(5 * time.Second)
	fmt.Println("主协程结束")
}

func worker(ctx context.Context, name string) {
	// 读取 context 中的值
	userId := ctx.Value("userId")
	role := ctx.Value("role")

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 停止, userId=%v, role=%v, 原因=%v\n",
				name, userId, role, ctx.Err())
			return
		default:
			fmt.Printf("%s: 工作中, userId=%v, role=%v\n", name, userId, role)
			time.Sleep(1 * time.Second)
		}
	}
}
