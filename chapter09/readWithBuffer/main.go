package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func GetFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "..", "hello.txt")
	return path
}
func main() {
	path := GetFilePath()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open-", err)
	}
	fmt.Printf("file=%v\n", file) //file=&{0xc000072788} 指针
	defer file.Close()
	reader := bufio.NewReader(file) //4096
	for {
		str, err := reader.ReadString('\n') //一行一行地读
		if err == io.EOF {
			break
		}
		fmt.Printf("%v\n", str)

	}
	fmt.Println("读完了...")
}
