package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func main() {
	path1 := GetFilePath(1)
	path2 := GetFilePath(2)
	data, err := ioutil.ReadFile(path1)
	if err != nil {
		fmt.Println("read file err=%v", err)
		return
	}
	err = ioutil.WriteFile(path2, data, 0666)
	if err != nil {
		fmt.Println("write file err=%v", err)
	}

}

func GetFilePath(num int) string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	var path string
	if num == 1 {
		path = filepath.Join(dir, "..", "zzz.txt")
	} else {
		path = filepath.Join(dir, "..", "yyy.txt")
	}
	return path
}
