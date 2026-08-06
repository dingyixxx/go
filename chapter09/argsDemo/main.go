package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数有：", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("os.Args[%v]=%v\n", i, v)
	}

	//go build -o test.exe calc_test.go
	//./test.exe tom d://testpath/access.log 909

	//命令行的参数有： 4
	//os.Args[0]=./test.exe
	//os.Args[1]=tom
	//os.Args[2]=d://testpath/access.log
	//os.Args[3]=909

	//windows电脑，就直接执行了
	//mac电脑./xxx.exe

	//用flag包来解析命令行参数

}
