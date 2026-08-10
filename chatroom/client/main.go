package main

import "fmt"

var userId int
var userPwd string

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
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	//更加用户的输入，显示新的提示信息
	if key == 1 {
		//说明用户要登陆
		fmt.Println("请输入用户的id")
		//fmt.Scanf("%d\n", &userId)
		fmt.Scanln(&userId)
		fmt.Println("请输入用户的密码")
		//fmt.Scanf("%v\n", &userPwd)
		fmt.Scanln(&userPwd)
		err := Login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}

	} else if key == 2 {

	}
}
