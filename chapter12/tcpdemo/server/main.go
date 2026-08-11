package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		fmt.Printf("服务器在等待客户端%s输入...\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Server read err...", err)
			return

		} else {
			fmt.Printf(string(buf[:n]) + "\n")
		}
	}
}

// net包
func main() {
	fmt.Println("Server starts listening...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Server failed to listen...", err)
		return
	}
	defer listen.Close() //closing
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("fail to connect...", err)
		} else {
			fmt.Println("Server is listening")
			fmt.Println("established...", conn.RemoteAddr())
		}
		go process(conn)
		//	这里准备起一个协程，服务客户端
	}
}
