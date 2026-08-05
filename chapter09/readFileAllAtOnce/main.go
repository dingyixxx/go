package main

import (
	"fmt"
	"io/ioutil"
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
	//ReadFile会自动打开和关闭, 因此, 无需显式地打开和关闭
	//这种"一次性读出来"的方案, 只适用于小文件
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	//file是字节数组
	fmt.Println(string(file))
}
