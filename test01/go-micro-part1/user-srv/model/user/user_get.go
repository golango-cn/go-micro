package user

import (
	log "github.com/micro/go-micro/v2/logger"
	"go-micro-part1/user-srv/basic/db"
	proto "go-micro-part1/user-srv/proto/user"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, user_pwd FROM _user WHERE user_name = ?`

	// 获取数据库
	o := db.GetDB()

	ret = &proto.User{}

	// 查询
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Error("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}
