package main

import (
	"context"
	"time"
	"github.com/coreos/etcd/clientv3"
	"fmt"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	if nil != err {
		fmt.Println("clientv3 New failed")
		return
	}
	defer cli.Close()

	// 可以测试超时退出：
	// ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/home/yangfan03", "homebrew")
	cancel()
	if nil != err {
		fmt.Printf("cli Put '/home/yangfan03' failed", err)
		return
	}

	readByEtcd("/home/yangfan03")
}

func readByEtcd(key string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	if nil != err {
		fmt.Println("clientv3.New failed")
		return
	}
	defer cli.Close()

	ctx, concel := context.WithTimeout(context.Background(), time.Second * 2)
	rsp, err    := cli.Get(ctx, key)
	concel()
	if nil != err {
		fmt.Println("cli Get failed")
		return
	}

	for _, ev := range rsp.Kvs {
		fmt.Println(string(ev.Key))
		fmt.Println(string(ev.Value))
		fmt.Println("--- --- --- --- ---")
	}
}