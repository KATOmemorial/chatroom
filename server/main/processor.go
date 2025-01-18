package main

import (
	"01/chatroom/common/message"
	"01/chatroom/server/processes"
	"01/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 功能：根据客户端发送的消息种类不同，决定调用哪个函数
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
				return err
			} else {
				fmt.Println("conn.Read err=", err)
				return err
			}
		}
		//fmt.Println("mes =", mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
