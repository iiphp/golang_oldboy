package main

import "fmt"

func main() {
	x := 0
	defer
		fmt.Println("First", x)
	x++
	defer fmt.Println("Secode:", x)
}
