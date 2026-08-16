package main

import (
	"errors"
	"fmt"
	"os"
)

// CircleQueue 使用结构体管理环形队列
type CircleQueue struct {
	maxSize int    // 队列最大容量
	array   [5]int // 底层数组
	head    int    // 指向队首
	tail    int    // 指向队尾
}

// AddQueue 向队列添加数据
func (this *CircleQueue) AddQueue(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// GetQueue 从队列获取数据
func (this *CircleQueue) GetQueue() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

// ListQueue 显示队列中的所有元素
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下:")
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}

	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head //永远空最后一格，好区分“空”和“满”（进而也可准确地衡量size）
}

func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	//	环形队列 redis主从复制 mysql的redoLog
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入1 表示添加数据到队列")
		fmt.Println("2. 输入2 表示从队列获取数据")
		fmt.Println("3. 输入3 表示显示队列")
		fmt.Println("4. 输入4 表示退出程序")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "2":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "3":
			queue.ListQueue()
		case "4":
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}

//	稀疏数组 压缩
//  环形链表 约瑟夫环 不写了

//goland两个不好的地方
//1.函数显示的2 usages，会把注释的单词也统计在内
//2.refactor“子文件夹”名字的时候，会把其他“大文件夹”的同名“子文件夹”里面的出现的同“包名”、打印出来的“文本”或者“注释”也给改了，非常不方便
