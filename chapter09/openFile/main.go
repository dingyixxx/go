package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	path := GetFilePath()
	//file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	//file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close() //defer捕获错误的话，也是一个栈
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}

	str := "双双， 双双\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 2; i++ {
		writer.WriteString(str)
	}
	writer.Flush() //和java一样的，写 + 反冲
}

func GetFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "..", "abc.txt")
	return path
}
