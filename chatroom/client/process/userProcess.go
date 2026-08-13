package process

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	utils "go_code/chatroom/utils"
	"net"
	"os"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//fmt.Printf("您输入的用户名是:%v,密码是:%v\n", userId, userPwd)
	//return nil
	//1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3. 创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 7. 到这个时候 data就是我们要发送的消息
	// 7.1 先把 data的长度发送给服务器
	// 先获取到 data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Println("客户端，发送消息的长度ok，为", buf, string(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//休眠20
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠了20..")
	// 这里还需要处理服务器端返回的消息.
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
		fmt.Println("500err-", err)
		err = errors.New(loginResMes.Error)
	} //汤姆凯特

	return

}

func (this *UserProcess) Register(userId int,
	userPwd string, userName string) (err error) {

	//1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3. 创建一个LoginMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4. 将loginMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	mes, err = tf.ReadPkg() // mes 就是

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成 LoginResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")

	} else {
		fmt.Println(registerResMes.Error)
		fmt.Println("500err-", err)
		err = errors.New(registerResMes.Error)
	}

	os.Exit(0)
	return

}
