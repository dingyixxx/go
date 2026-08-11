package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	read, err := conn.Read(buf[:4])
	fmt.Println("读取客户端发送的数据..", read) //从conn读取字节到buf中

	if err != nil {
		fmt.Println("conn.Read err=", err)
		err = errors.New("读长度出错了")
		return
	}
	//根据buf[:4] 转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	fmt.Println("读到的buf=", buf[:4])

	//根据pkgLen读取消息内容
	//从conn读pkgLen个字节,扔到buf
	n, err := conn.Read(buf[:pkgLen])
	fmt.Println("n=", n)
	if n != int(pkgLen) || err != nil {
		err = errors.New("读内容出错了")
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//循环的客户端发送的信息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务器端也跳出循环了")
				return
			}
			fmt.Println("readPkg err=", err)
			return
		}
		fmt.Println("mes=", mes)
	}
}

func main() {

	//提示信息
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端来链接服务器.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通信
		go process(conn)
	}
}
