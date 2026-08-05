package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	path := GetFilePath()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open-", err)
	}
	fmt.Printf("file=%v", file) //file=&{0xc000072788} 指针

	err = file.Close()
	if err != nil {
		fmt.Println("close-", err)
	}
}

func GetFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "..", "hello.txt")
	return path
}
