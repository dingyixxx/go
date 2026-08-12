package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

// 编写一个函数serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
