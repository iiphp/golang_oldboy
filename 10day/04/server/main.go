package main

import (
	"net"
	"fmt"
	"io"
	"time"
	"github.com/gomodule/redigo/redis"
	"golang_oldboy/10day/04/model"
)

var (
	pool    *redis.Pool
	userMgr *model.UserMgr
)

func main()  {
	pool = initRedis("localhost:6379", 8, 32, time.Second*300)
	userMgr = initUserMgr(pool)
	doServer("0.0.0.0:8888")
}

func initUserMgr(pool *redis.Pool) (*model.UserMgr) {
	mgr := model.NewUserMgr(pool)
	return mgr
}

func doServer(addr string) (err error) {
	listen, err := net.Listen("tcp", addr)
	if nil != err {
		fmt.Println("net.Listen failed", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if nil != err {
 			fmt.Println("listen.Accept failed", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn)  {
	// 上来先 defer 资源处理
	defer func() {
		conn.Close()
	}()

	// 把 conn 封装为一个 client
	client := newClient(conn)
	fmt.Println("toProcess")
	if err := client.process(); nil != err {
		if io.EOF == err {
			fmt.Println("client.process client close")
			return
		}
		fmt.Println("client.process failed", err)
	}
}
