package processes

import "fmt"

var (
	userMgr *UerMgr
)

type UerMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UerMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (this *UerMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

func (this *UerMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

func (this *UerMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

func (this *UerMgr) GetOnlineUsrById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
