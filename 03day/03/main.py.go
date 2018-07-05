package main

import "fmt"

// 函数也可以作为一种类型被定义，被赋值

type add_func func(int, int) int

func main() {
	var f add_func = add
	var x = f(1, 20)
	fmt.Println(x)
}

func add(a int, b int) int {
	return a + b
}

