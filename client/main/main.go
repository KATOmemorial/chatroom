package main

import (
	"01/chatroom/client/process"
	"fmt"
)

var userId int
var userPwd string
var userName string

func main() {
	//接收用户选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true

	for {
		fmt.Println("-----------------欢迎登录-----------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%d\n", &userPwd)
			fmt.Println("请输入用户名字:")
			fmt.Scanf("%d\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
		default:
			fmt.Println("输入有误")
		}
	}

	/*
		//根据用户输入，显示对应的信息
		if key == 1 {
			//说明用户要登录
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			login(userId, userPwd)
			//if err != nil {
			//	fmt.Println("登录失败")
			//} else {
			//	fmt.Println("登录成功")
			//}
		} else if key == 2 {
			fmt.Println("jin")
		}
	*/
}
