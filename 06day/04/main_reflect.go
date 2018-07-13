package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string
	Age int "json:age"
	Score int
}

func foo(input interface{})  {
	v := reflect.ValueOf(input)
	if v.Elem().Kind() != reflect.Struct {
		fmt.Println("input is not a struct")
		return
	}

	fmt.Println("NumFiled is", v.Elem().NumField())
	fmt.Println("NumMethod is", v.Elem().NumMethod())
}

// 涉及 reflect 的各种库函数，比如 Field，Tag 等等
func main() {
	std := Student{Name:"Tom", Age:17, Score:98}
	foo(&std)
}
