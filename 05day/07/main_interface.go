package main

import "fmt"

type Play interface {
	Ready()
	Run()
}

type Student struct {
	Name string
	Age int
}

func (p *Student) Ready() {
	fmt.Println("Ready Ready")
}

func (p *Student) Run() {
	fmt.Println("Run Run Run")
}

func main() {
	var play Play
	std := Student{Name:"Tom", Age:16}
	play = &std
	play.Ready()
	play.Run()
}
