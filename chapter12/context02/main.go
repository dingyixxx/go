package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. 创建可手动取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 2. 绑定值
	ctx = context.WithValue(ctx, "taskId", "task-001")
	ctx = context.WithValue(ctx, "priority", "high")

	go worker(ctx, "协程1")
	go worker(ctx, "协程2")
	go worker(ctx, "协程3")

	// 3. 5秒后手动取消
	time.Sleep(5 * time.Second)
	fmt.Println("主协程决定取消所有任务...")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("主协程结束")
}

func worker(ctx context.Context, name string) {
	taskId := ctx.Value("taskId")
	priority := ctx.Value("priority")

	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 停止, taskId=%v, priority=%v, 原因=%v\n",
				name, taskId, priority, ctx.Err())
			return
		default:
			fmt.Printf("%s: 第%d次工作, taskId=%v\n", name, i, taskId)
			time.Sleep(1 * time.Second)
		}
	}
}
