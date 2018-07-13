package main

import "fmt"

// 改为 只读/只写 channel 模式： func send(inputChan chan<- int, exitChan chan<- bool) {
func send(inputChan chan int, exitChan chan bool) {
	for i := 0; i < 20; i++ {
		inputChan <- i
	}
	close(inputChan)
	exitChan <- true
}

// 改为 只读/只写 channel 模式： func recv(inputChan <-chan int, exitChan chan<- bool) {
func recv(inputChan chan int, exitChan chan bool) {
	for v := range inputChan {
		fmt.Println(v)
	}
	exitChan <- true
}

func main() {
	exitChan := make(chan bool, 2)
	inputChan := make(chan int, 10)
	go send(inputChan, exitChan)
	go recv(inputChan, exitChan)

	for i := 0; i < 2; i++ {
		<- exitChan
	}
}
