package main

import (
	"time"
	"fmt"
)

func toPanic() {
	defer func() {
		if err := recover(); nil != err {
			fmt.Println("toPanic is here")
		}
	}()

	var m map[int]string
	m[1] = "abc"
}

func main()  {
	go toPanic()
	time.Sleep(time.Second)
}
