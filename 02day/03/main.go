package main

import (
	"time"
	"fmt"
)

const (
	Male   = 1
	Female = 2
)

func main() {
	t := time.Now().Unix()
	fmt.Println(t)

	if (0 == t % Female) {
		fmt.Println("Male")
	} else {
		fmt.Println("Female")
	}
}
