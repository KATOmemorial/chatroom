package process

import (
	"01/chatroom/client/utils"
	"01/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userId = %d userPwd=%s", userID, userPwd)
	//return nil

	//链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//创建以后LoginMes实例
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//把data赋给mes.data字段
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	fmt.Printf("客户端发送消息的长度=%d 内容是%s", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//休眠20
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠了20..")

	//这里还需要处理服务器端返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		fmt.Println("当前在线用户列表：")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("\n\n")
		go ProcessServerMes(conn)
		//fmt.Println("success")
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//创建以后LoginMes实例
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//把data赋给mes.data字段
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritrPkg err=", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	var registerResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请前往登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
