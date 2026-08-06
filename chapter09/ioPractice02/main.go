package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

var count CharCount

func GetFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	var path string
	path = filepath.Join(dir, "..", "statistics.txt")
	return path
}

func main() {
	filePath := GetFilePath()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, value := range str {
			switch {
			case value >= 'a' && value <= 'z':

				fallthrough
			case value >= 'A' && value <= 'Z':

				count.ChCount++
			case value == ' ' || value == '\t':
				count.SpaceCount++
			case value >= '0' && value <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字母是:%v，数字是:%v，空格是:%v，其他字符是:%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}

//type CharCount struct {
//	ChCount    int
//	NumCount   int
//	SpaceCount int
//	OtherCount int
//}
