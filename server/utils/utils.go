package utils

import (
	"01/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	//fmt.Println("读到的buf=", buf[:4])
	pkgLen := binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	//把buf反序列化
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	pkgLen := uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
