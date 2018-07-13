package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	for {
		select {
		case v := <- intChan:
			fmt.Println(v)
		default:
			fmt.Println("i am here")
			time.Sleep(time.Second)
		}
	}
}
