package main

import "fmt"

func w(inputChan chan int, outputChan chan int, exitChan chan bool) {
	for v := range inputChan {
		if v % 2 == 0 {
			outputChan <- v
		}
	}

	flag := <- exitChan
	if !flag {
		close(outputChan)
	}
}


func main()  {
	max := 10
	input := make(chan int, 100)
	output := make(chan int, 100)
	exitChan := make(chan bool, 3)

	// 跳过，开了 3 个 goroutine
	for i := 0; i < 3; i++ {
		if i == 2 {
			exitChan <- false
		} else {
			exitChan <- true
		}

		go w(input, output, exitChan);
	}
	close(exitChan)

	for i := 0; i < max; i++ {
		input <- i
	}
	close(input)

	for v := range output {
		fmt.Println(v)
	}
}
