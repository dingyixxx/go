package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn   net.Conn
	UserId int
}

//ServerProcessRegister

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes 中取出 mes.Data，并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//1先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterMesType
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库去完成注册.
	//1.使用model.MyUserDao 到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
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
	//6. 发送data，我们将其封装到writePkg函数
	//因为使用分层模式(mvc)，我们先创建一个Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

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

	//我们需要到redis数据库去完成验证.
	//1.使用model.MyUserDao 到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		//loginResMes.Code = 500
		//loginResMes.Error = "该用户不存在，请注册再使用..."
		//这里我们先测试成功!,然后我们在可以根据返回具体错误信息
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//将当前在线用户的id 放入到loginResMes.UsersId
		//遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserId = append(loginResMes.UserId, id)
		}
		fmt.Println(user, "登录成功")
	}

	//如果用户id= 100， 密码=123456，认为合法，否则不合法
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	loginResMes.Code = 200
	//	loginResMes.Error = "用户登录成功"
	//	//合法
	//} else {
	//	//不合法
	//	loginResMes.Code = 500
	//	loginResMes.Error = "表示该用户不存在"
	//}

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
