package main

import "fmt"

type Score struct {
	Englist int
	Math int
}

type Student struct {
	Name string
	Age int
	Score
	int
}

func main() {
	var std Student
	std.Name = "Tom"
	std.Age = 21
	std.Englist = 98
	std.Math = 66
	std.int = 100

	fmt.Println(std)

	// 这种写法也可以，我感觉这样写更清楚
	std.Score.Englist += 20
	std.Score.Math += 10

	fmt.Println(std)
}
