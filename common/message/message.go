package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的内容
}

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code    int `json:"code"` //返回状态码 500 表示该用户未注册 200 表示登录成功
	UsersId []int
	Error   string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码 400 表示该用户已经占用 200 表示注册成功
	Error string `json:"error"` //返回错误信息
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMes struct {
	Content string `json:"content"`
	User
}
