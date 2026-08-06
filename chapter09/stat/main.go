package main

import (
	"os"
	"path/filepath"
	"runtime"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	path := GetFilePath(1)
	exists, err := PathExists(path)
	if err == nil {
		println(exists)
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
