package main

import (
	"net"
	"golang_oldboy/10day/04/model"
	"golang_oldboy/10day/04/proto"
	"fmt"
	"encoding/binary"
	"encoding/json"
	"io"
)

type client struct {
	conn net.Conn
	buf  [8192]byte
}

func newClient(conn net.Conn) *client {
	cli := &client{
		conn:conn,
	}
	return cli
}

func (this *client) process() (err error) {
	var msg proto.Message
	i := 0
	// 因为聊天室是长连接嘛，所以这里是「死循环」
	for {
		i++
		fmt.Println("process =", i)
		msg, err = this.readPackage()
		// 如果读消息失败了
		if nil != err {
			if io.EOF == err {
				break
			}
			fmt.Println("readPackage failed")
			return
		}

		err = this.processMessage(msg)
		if nil != err {
			fmt.Println("processMessage failed")
			continue
		}
	}
	return
}

func (this *client) readPackage() (msg proto.Message, err error) {
	// 先把消息的长度读出来
	n, err := this.conn.Read(this.buf[0:4])

	// 如果客户端关闭了连接
	if io.EOF == err {
		return
	}

	if nil != err || 4 != n {
		fmt.Println("this.conn.Read packLen failed", n, err)
		return
	}
	fmt.Println("readPackage:", this.buf[0:4])


	var packLen uint32
	packLen = binary.BigEndian.Uint32(this.buf[0:4])
	n, err = this.conn.Read(this.buf[0:packLen])
	if nil != err || n != int(packLen) {
		fmt.Println("this.conn.Read content failed")
		return
	}

	fmt.Printf("receive data:%s\n", string(this.buf[0:packLen]))
	err = json.Unmarshal(this.buf[0:packLen], &msg)
	if nil != err {
		fmt.Println("json.Unmarshal failed")
		return
	}
	return
}

func (this *client) writePackage(jsonMsgData []byte) (err error)  {
	// 写消息长度
	packLen := uint32(len(jsonMsgData))
	binary.BigEndian.PutUint32(this.buf[0:4], packLen)

	n, err := this.conn.Write(this.buf[0:4])
	if nil != err || 4 != n {
		fmt.Println("conn.write packLen failed")
		return
	}

	// 写消息内容
	n, err = this.conn.Write(jsonMsgData)
	if nil != err || int(packLen) != n {
		fmt.Println("conn.Write connect failed")
		return
	}
	return
}

func (this *client) processMessage(msg proto.Message) (err error) {
	switch msg.Cmd {
	case proto.CmdLoginRequest:
		err = this.login(msg)
	case proto.CmdRegRequest:
		err = this.register(msg)
	default:
		err = model.ErrMethodIsNotExist
	}
	return
}

func (this *client) login(msg proto.Message) (err error) {
	defer func() {
		this.loginResponse(err)
	}()

	var user model.User
	user = model.User{Id:101, Passwd:"3g3w", Name:"Tom", Age:17}

	var loginCmd proto.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &loginCmd)
	if nil != err {
		fmt.Println("json.Unmarshal loginCmd failed")
		return
	}

	if loginCmd.Passwd != user.Passwd {
		err = model.ErrInvalidPasswd
	}
	return
}

func (this *client) loginResponse(err error) {
	// 准备消息结构体
	msg := proto.Message{Cmd:proto.CmdLoginResponse, Data:""}
	loginRsp := proto.LoginRsp{Code:200, Error:""}
	if nil != err {
		loginRsp.Code = 500
		loginRsp.Error = fmt.Sprintf("%s", err)
	}

	jsonLoginRsp, err := json.Marshal(loginRsp)
	if nil != err {
		fmt.Println("json.Marshal loginRsp failed")
		return
	}

	msg.Data = string(jsonLoginRsp)
	jsonMsgData, err := json.Marshal(msg)
	if nil != err {
		fmt.Println("json.Marshal msg failed")
		return
	}

	err = this.writePackage(jsonMsgData)
	if nil != err {
		fmt.Println("writePackage jsonMsgData failed")
		return
	}
	return
}

func (this *client) register(msg proto.Message) (err error)  {
	defer func () {
		this.registerResponse(err)
	}()

	var regCmd proto.RegCmd
	err = json.Unmarshal([]byte(msg.Data), &regCmd)
	if nil != err {
		fmt.Println("register json.Unmarshal regCmd failed", err)
		return
	}

	fmt.Println("register receive user is:", regCmd.User)
	err = userMgr.Register(&regCmd.User)

	// Todo 记录到 Redis
	return
}

func (this *client) registerResponse(err error) {
	// 构建消息
	msg := proto.Message{Cmd:proto.CmdRegResponse, Data:""}
	regRsp := proto.RegRsp{Code:200, Error:""}
	if nil != err {
		regRsp.Code  = 500
		regRsp.Error = fmt.Sprintf("%s", err)
	}

	jsonRegRsp, err := json.Marshal(regRsp)
	if nil != err {
		fmt.Println("registerResponse json.Marshal regRsp failed")
		return
	}

	msg.Data = string(jsonRegRsp)
	jsonMsgData, err := json.Marshal(msg)
	if nil != err {
		fmt.Println("registerResponse json.Marshal msg failed")
		return
	}

	err = this.writePackage(jsonMsgData)
	if nil != err {
		fmt.Println("registerResponse writePackage jsonMsgData failed")
		return
	}

	return
}