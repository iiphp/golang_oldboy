package main

import "fmt"

func main()  {
	intChan := make(chan int, 10)
	for i := 0; i < 8; i++ {
		intChan <- i
	}

	close(intChan)

	for {
		r, ok := <- intChan

		// 判断 channel 被关闭
		if !ok {
			break
		}
		fmt.Println(r)
	}
}
