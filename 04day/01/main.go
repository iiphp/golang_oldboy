package main

import "fmt"

func main() {
	var x []int
	fmt.Println("x init: ", x)
	x = append(x, 2, 3, 4)
	fmt.Println("x, append: ", x)
	x = append(x, x...)
	fmt.Println("x, append self: ", x)
}