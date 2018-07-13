package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

func main()  {
	stdChan := make(chan interface{}, 10)
	std := Student{Name:"Tom", Age:15}
	stdChan <- &std

	std1 := <- stdChan
	fmt.Println(reflect.TypeOf(std1), reflect.ValueOf(std1).Kind(), std1)
	std2, ok := std1.(*Student)
	if !ok {
		fmt.Println("Interface transform fail")
	}
	fmt.Println(reflect.TypeOf(std2), reflect.ValueOf(std2).Kind(), std2)
}
