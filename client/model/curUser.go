package model

import (
	"01/chatroom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
