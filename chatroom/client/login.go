package main

import "fmt"

func Login(userId int, userPwd string) (err error) {
	fmt.Printf("您输入的用户名是:%v,密码是:%v\n", userId, userPwd)
	return nil
}
