package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p *Person) Eat(food string) {
	fmt.Println("eat", food)
}

type Student struct {
	Person
	School string
}

func main() {
	//var std Student
	//std.Name = "Tom"
	//std.Age = 17
	//std.School = "Fuxue"

	std := Student{Person{Name:"Tom", Age:17}, "Fuxue"}

	fmt.Println(std)
	std.Eat("banana")
}
