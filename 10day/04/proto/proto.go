package proto

import "golang_oldboy/10day/04/model"

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type LoginRsp struct {
	Code int     `json:code`
	Error string `json:error`
}

type RegCmd struct {
	User model.User `json:user`
}

type RegRsp struct {
	Code  int
	Error string
}