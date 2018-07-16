package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
)

func main()  {
	filename := "/tmp/tail.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset:0, Whence:2},
		MustExist: false,
		Poll:      true,
	})

	if nil != err {
		fmt.Println("tail.TailFile failed")
		return
	}

	for {
		msg, ok := <- tails.Lines

		// 如果管道关闭
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Millisecond * 100)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
