package main

import (
	"time"
	"fmt"
)

func main() {
	n := 3
	l := 4
	c := make(chan int, n)
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key%d", i)
		go doSth(key, l, c)
	}

	for i := 0; i < n * l; i++ {
		<- c
	}
}

func doSth(key string, l int, c chan int) {
	for i := 0; i < l; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("[", key, "] [", i, "] doSth")
		c <- i
	}
}