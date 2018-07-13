package main

import (
	"net"
	"fmt"
	"encoding/json"
	"golang_oldboy/10day/04/proto"
	"encoding/binary"
	"golang_oldboy/10day/04/model"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:8888")
	defer func() {
		conn.Close()
	}()
	if nil != err {
		fmt.Println("net.Dial failed")
		return
	}

	login(conn, "Tom", "memo")
	login(conn, "Tom", "3g3w")
	register(conn, 113, "u3z4", "Jerry", 17)
}

func login(conn net.Conn, userName string, passwd string) (err error) {

	// 拼消息
	msg := proto.Message{Cmd:proto.CmdLoginRequest, Data:""}

	loginCmd := proto.LoginCmd{Name:userName, Passwd:passwd}
	jsonData, err := json.Marshal(loginCmd)
	if nil != err {
		fmt.Println("json.Marshal(loginCmd) failed")
		return
	}
	msg.Data = string(jsonData) // Message 结构体准备好了

	err = writePackage(conn, msg)
	if nil != err {
		fmt.Println("register writePackage failed", err)
	}

	// 读返回结果，并打印
	msg, err = readPackage(conn)
	if nil != err {
		fmt.Println("readPackage failed", err)
		return
	}

	fmt.Println(msg)
	return
}

func register(conn net.Conn, id int, passwd string, name string, age int)  {
	// 准备消息
	msg := proto.Message{Cmd:proto.CmdRegRequest, Data:""}
	regCmd := proto.RegCmd{User:model.User{}}
	regCmd.User = model.User{Id:id, Passwd:passwd, Name:name, Age:age}
	jsonRegCmd, err := json.Marshal(regCmd)
	if nil != err {
		fmt.Println("json.Marshal user failed")
		return
	}
	msg.Data = string(jsonRegCmd)

	err = writePackage(conn, msg)
	if nil != err {
		fmt.Println("register writePackage failed", err)
	}

	// 读返回结果，并打印
	msg, err = readPackage(conn)
	if nil != err {
		fmt.Println("register readPackage failed", err)
		return
	}
	fmt.Println(msg)
}

func writePackage(conn net.Conn, msg proto.Message) (err error) {
	jsonMsg, err := json.Marshal(msg)
	if nil != err {
		fmt.Println("json.Marshal msg failed")
		return
	}

	// 写消息长度
	var buf [4]byte
	packLen := uint32(len(jsonMsg))
	binary.BigEndian.PutUint32(buf[0:4], packLen)
	n, err := conn.Write(buf[0:4])
	if nil != err || 4 != n {
		fmt.Println("conn.Write packLen failed", msg, n, err)
		return
	}

	// 写消息内容
	n, err = conn.Write(jsonMsg)
	if nil != err || int(packLen) != n {
		fmt.Println("conn.Write jsonMsg failed", msg, n, err)
		return
	}
	return
}

func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var readData [8192]byte
	n, err := conn.Read(readData[0:4])
	if nil != err || 4 != n {
		fmt.Println("conn.Read packLen failed", n, err)
		return
	}

	var packLen uint32
	packLen = binary.BigEndian.Uint32(readData[0:4])

	n, err = conn.Read(readData[0:packLen])
	if nil != err || int(packLen) != n {
		fmt.Println("conn Read content failed")
		return
	}

	err = json.Unmarshal(readData[0:packLen], &msg)
	if nil != err {
		fmt.Println("json.Unmarshal msg failed")
		return
	}
	return
}