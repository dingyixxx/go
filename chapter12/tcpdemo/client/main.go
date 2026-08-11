package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Client failed to connect...", err)
		return
	}
	fmt.Println("established...", conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		content, err1 := reader.ReadString('\n')
		if err1 != nil {
			fmt.Println("reader.ReadString failed...", err1)
		}
		content = strings.TrimRight(content, "\n\r")
		if content == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		n, err2 := conn.Write([]byte(content))
		if err2 != nil {
			fmt.Println(" conn.Write([]byte(content)) failed...", err2)
		}
		fmt.Printf("客户端发送了%d个字节\n", n)
	}

}
