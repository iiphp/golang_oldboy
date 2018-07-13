package model

import "errors"

var (
	ErrInvalidPasswd    = errors.New("无效密码")
	ErrNameIsNotExist   = errors.New("用户不存在")
	ErrNameIsExist      = errors.New("用户已存在")
	ErrMethodIsNotExist = errors.New("方法不存在")
)
