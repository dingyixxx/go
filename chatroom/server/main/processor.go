package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	process2 "go_code/chatroom/server/process"
	"go_code/chatroom/utils"

	"io"
	"net"
)

// 先创建一个Processor 的结构体体
type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMes 函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	fmt.Println("mes=", *mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理登录登录
		//处理登录登录
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)

	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)

	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {
	//循环的客户端发送的信息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		var mes message.Message
		mes, err = tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
				return err
			}

			fmt.Println("readPkg err=", err)
			return err
		}

		err = this.serverProcessMes(&mes)
		if err != nil {
			return
		}
		fmt.Println("mes=", mes)
	}
}
