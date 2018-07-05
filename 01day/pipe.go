package main

import "fmt"

func pipe() {
	p := make(chan int, 3)
	p <- 10
	p <- 20

	fmt.Println(len(p))

	var tmp int
	tmp = <- p

	fmt.Println(tmp)
	fmt.Println(len(p))
}