package main

import (
	"net"
	"fmt"
	"io"
)

func main()  {
	 doServer("0.0.0.0:8888")
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
