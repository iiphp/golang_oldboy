package main

import "fmt"

// 可变参数的函数

func main() {
	var x = f0(1, 20, 300)
	fmt.Println(x)
}


// 0 个或多个参数
func f0(arg ...int) int {
	sum := 0
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

// 1 个或多个参数
func f1(a int, arg ...int) int {
	return 0
}

// 2 个或多个参数
func f2 (a int, b int, arg ...int) int {
	return 0
}