package main

import (
	"fmt"
	"os"
)

func isExeByMagic(filePath string) (bool, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer f.Close()

	buf := make([]byte, 2)
	_, err = f.Read(buf)
	if err != nil {
		return false, err
	}

	// PE 文件以 "MZ" 开头
	return buf[0] == 'M' && buf[1] == 'Z', nil
}

func main() {
	ok, err := isExeByMagic("/Users/bytedance/Desktop/go/goproject/hello")
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	fmt.Println("是exe文件:", ok)
}
