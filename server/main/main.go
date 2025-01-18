package main

import (
	"01/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	processor.process2()
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器端通讯错误 err=", err)
		return
	}
}

func initUseDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUseDao()
	fmt.Println("服务器【新结构】在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端来链接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		go process(conn)
	}
}
