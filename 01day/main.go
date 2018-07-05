package main

import (
	"fmt"
)

func main ()  {
	fmt.Println("Hello World")

	fmt.Println(add(1, 30))

	pipe()
}

func add (a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

