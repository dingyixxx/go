package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func CopyFile(dstFilePath string, srcFilePath string) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return 0, err
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)

}

func main() {
	path := GetFilePath(1)
	path2 := GetFilePath(2)
	CopyFile(path2, path)

}
func GetFilePath(num int) string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	var path string
	if num == 1 {
		path = filepath.Join(dir, "..", "li.webp")
	} else {
		path = filepath.Join(dir, "..", "li2.webp")
	}
	return path
}
