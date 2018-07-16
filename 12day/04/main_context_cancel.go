package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	foo()
	time.Sleep(time.Second)
}

func foo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 在本函数退出时候，调用 cancel

	c := gen(ctx)
	for n := range c {
		fmt.Println(n)
		if 4 == n {
			break
		}
	}
}


// 如果 ctx 不调用 cancel ，就一直运行 num++
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	num := 1
	go func() {
		for {
			select {
			// 如果取消
			case <-ctx.Done():
				fmt.Println("gen will exit")
				return
			case dst <- num:
				num++
			}
		}
	}()
	return dst
}