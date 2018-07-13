package main

import (
	"time"
	"fmt"
)

func main() {

	for {
		// 每个 1 秒执行一次
		t := time.NewTicker(time.Second)
		select {
			case <- t.C:
				fmt.Println("hello")
		}
		t.Stop()
	}
}
