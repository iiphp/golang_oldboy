package main

import (
	"golang_oldboy/02day/02/add"
	"fmt"
)

func init() {
	fmt.Println("There is main package's init func")
}

func main ()  {
	fmt.Println("There is main func")
	fmt.Println("add.Name is " + add.Name)
}