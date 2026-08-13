package main

import (
	"fmt"
	"go_code/chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	//判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("-----------------欢迎登陆多人聊天系统-----------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		//fmt.Scanf("%d\n", &key)
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			//fmt.Scanf("%d\n", &userId)
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码")
			//fmt.Scanf("%v\n", &userPwd)
			fmt.Scanln(&userPwd)
			processor := &process.UserProcess{}
			err := processor.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登录失败")
			} else {
				fmt.Println("登录成功")
			}
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)
			processor := &process.UserProcess{}
			err := processor.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("注册失败")
			} else {
				fmt.Println("注册成功")
			}
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	////更加用户的输入，显示新的提示信息
	//if key == 1 {
	//	//说明用户要登陆
	//

	//
	//} else if key == 2 {
	//
	//}
}
