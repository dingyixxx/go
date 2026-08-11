package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 作业 2：goroutine + channel 配合完成排序，并写入文件
// 要求：
// 1）开 10 个协程 writeDataToFile，每个协程随机生成 1000 个数据，存放到 10 文件中
// 2）当 10 个文件都生成了，让 10 个 sort 协程从 10 文件中读取 1000 个文件，并完成排序，重新写入到 10 个结果文件
var count int = 10
var thousand int = 1000
var ceil int = 99000
var writeChan chan int = make(chan int, count)
var readChan chan int = make(chan int, count)
var mu sync.Mutex
var entireSlice []int = make([]int, 0)
var rewriteChan chan int = make(chan int, count)

func main() {
	for i := 0; i < count; i++ {
		go writeDataToFile(i)
	}
	for i := 0; i < count; i++ {
		<-writeChan
	}
	fmt.Println("10个协程都写完了...")
	for i := 0; i < count; i++ {
		go readDataToMemory(i)
	}
	for i := 0; i < count; i++ {
		<-readChan
	}
	fmt.Printf("排序前---切片长度%v---%v", len(entireSlice), entireSlice)
	fmt.Println()
	sort.Ints(entireSlice)
	fmt.Println("10个协程都读完+排序了...")
	//更好地做法是"合并n个升序链表", 但是假设内存也够, 就不写了
	fmt.Printf("排序后---切片长度%v---%v", len(entireSlice), entireSlice)
	for i := 0; i < count; i++ {
		go rewriteDataToFile(i)
	}
	for i := 0; i < count; i++ {
		<-rewriteChan
	}
	fmt.Println("10个协程都重新写完了...")

}

func rewriteDataToFile(i int) {
	res := entireSlice[thousand*i : thousand*(i+1)]
	str := ""
	for j, v := range res {
		if j == len(res)-1 {
			str += strconv.Itoa(v)
		} else {
			str += strconv.Itoa(v) + ","
		}
	}
	fmt.Printf("我是rewrite协程%v,现在,我开始数了\n", i)
	//ioutil.ReadAll()
	//io.ReadFull() 能定容定容，避免多次扩容
	ioutil.WriteFile(GetFilePath(i, "final00"), []byte(str), 0666)
	writeChan <- 1 //这种, 我感觉也是最好写在最后一行
	fmt.Printf("我是rewrite协程%v,重写完了\n", i)
	time.Sleep(time.Second * time.Duration(rand.Intn(10)+1))
	rewriteChan <- 1
}

func readDataToMemory(i int) []int {
	var res []int = make([]int, 0)
	fmt.Printf("我是read协程%v,现在,我开始数了\n", i)
	file, err := ioutil.ReadFile(GetFilePath(i, "temp00"))
	if err != nil {
		fmt.Println(err)
	}

	s := string(file)
	split := strings.Split(s, ",")
	for _, value := range split {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("文件temp00%v.txt的数值%v在strconv.Atoi(value)环节出错了", i, num)
			continue
		}
		res = append(res, num)
	}
	fmt.Printf("我是协程%v,文件temp00%v.txt读完了,切片长度%v\n", i, i, len(res))
	mu.Lock()
	entireSlice = append(entireSlice, res...)
	mu.Unlock()
	time.Sleep(time.Second * time.Duration(rand.Intn(10)+1))
	readChan <- 1
	return res
}

func writeDataToFile(i int) {
	fmt.Printf("我是协程%v,现在,我开始生成[0,%v)之间的%v个随机数,然后存放到temp00%v.txt文件里了\n", i, ceil, thousand, i)
	str := ""
	for i := 0; i < thousand; i++ {
		if i == thousand-1 {
			str += strconv.Itoa(rand.Intn(ceil))
		} else {
			str += (strconv.Itoa(rand.Intn(ceil)) + ",")
		}
	}
	time.Sleep(time.Second * time.Duration(rand.Intn(10)+1))
	var data []byte = []byte(str)
	ioutil.WriteFile(GetFilePath(i, "temp00"), data, 0666)
	writeChan <- 1 //这种, 我感觉也是最好写在最后一行
	fmt.Printf("我是协程%v,写完了\n", i)

}
func GetFilePath(num int, prefix string) string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "..", prefix+strconv.Itoa(num)+".txt")
}
