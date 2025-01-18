package process

import (
	"01/chatroom/client/model"
	"01/chatroom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

func updataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}
