package main

import "fmt"

// 结构体定义
type Student struct {
	Name string
	Age int
}

// 结构体初始化
func main() {
	var tom Student
	tom.Name = "Tom"
	tom.Age = 20
	fmt.Println(tom)

	var jerry *Student = &Student{
		Name: "Jerry",
		Age: 20,
	}
	fmt.Println(jerry)
}
