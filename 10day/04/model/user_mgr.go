package model

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"time"
)

type UserMgr struct {
	pool *redis.Pool
}

func NewUserMgr(redisPool *redis.Pool) *UserMgr {
	mgr := &UserMgr{
		pool:redisPool,
	}
	return mgr
}

func (this *UserMgr) Login(name string, passwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	err = nil

 	// Todo: 给 User 连接起来
 	user = &User{Id:101, Passwd:"abc", Name:"Tom", Age:15}

 	// 判断「密码」是否正确
 	if user.Passwd != passwd {
		err = ErrInvalidPasswd
		return
	}

	user.Status = UserStatusOnline
	user.LastLoginTime = fmt.Sprintf("%v", time.Now())
	return
}