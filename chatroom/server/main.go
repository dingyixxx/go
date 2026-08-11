package main

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/utils"
	"io"
	"net"
)

// 编写一个函数serverProcessLogin函数，专门处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码...
	//1. 先从mes 中取出 mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//在声明一个 LoginResMes
	var loginResMes message.LoginResMes
	//如果用户id= 100， 密码=123456，认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
		loginResMes.Error = "用户登录成功"
		//合法
	} else {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "表示该用户不存在"
	}

	//3将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4. 将data 赋值给 resMes
	resMes.Data = string(data)

	//5. 对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	err = utils.WritePkg(conn, data)
	return
}

// 编写一个ServerProcessMes 函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录登录
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//循环的客户端发送的信息
	for {
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务器端也跳出循环了")
				return
			}
			fmt.Println("readPkg err=", err)
			return
		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
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
