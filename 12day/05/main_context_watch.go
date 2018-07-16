package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func main() {
	go watchConf()
	time.Sleep(time.Second * 10)
}

// 测试时，在命令行用 etcdctl put /home/yangfan03 "devspace1" 修改
func watchConf() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout: time.Second * 3,
	})
	if nil != err {
		fmt.Println("clientv3.New failed")
		return
	}
	defer cli.Close()

	for {
		rch := cli.Watch(context.Background(), "/home/yangfan03")
		for wrsp := range rch {
			for _, ev := range wrsp.Events {
				fmt.Println(ev.Type, ev.Kv)
			}
		}
	}
}