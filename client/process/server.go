package process

import (
	"01/chatroom/client/utils"
	"01/chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("-------登录成功--------")
	fmt.Println("1.显示在线用户列表")
	fmt.Println("2.发送消息")
	fmt.Println("3.消息列表")
	fmt.Println("4.退出系统")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	fmt.Scanf("%d\n", &key)
	smsProcess := &SmsProcess{}
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入群发消息")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入不对")
	}
}

func ProcessServerMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//fmt.Printf("mes=%v", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updataUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回了一个未知消息类型")
		}
	}
}
