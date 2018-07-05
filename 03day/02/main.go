package main

import "fmt"

// 指针
func main() {
	var x int = 10
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println("--- --- --- ---")
	var y *int = &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(&y)
}
