package main

import "fmt"

type Student struct {
	Name string
	age int
}

func (p Student) getAge() int {
	return p.age
}

// 注意，是 *Student
// 因为，setAge 前面的 p 是结构体，默认是值传递
// 所以，如果不加 * 的话，setAge 里面改变，不会影响 main 中的 std 变量
func (p *Student) setAge(age int) {
	p.age = age
}

func main() {
	std := Student{Name:"Tom", age:21}
	fmt.Println(std.getAge())


	// 正确的写法是，(&std).setAge(42)，Go 为了方便，帮助我们做了一个只能的转化，所以可以写成下面这种简单的形式
	std.setAge(39)
	fmt.Println(std.getAge())
}
