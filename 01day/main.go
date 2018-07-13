package main

import (
	"fmt"
	"errors"
	"reflect"
)

func main ()  {
	fmt.Println("Hello World")

	fmt.Println(add(1, 30))

	pipe()

	err := errors.New("i am here")
	str1 := fmt.Sprintf("%v", err)
	str2 := fmt.Sprintf("%s", err)
	fmt.Println("--- --- str1 --- ---")
	fmt.Println(str1)
	fmt.Println(reflect.TypeOf(str1))
	fmt.Println(reflect.ValueOf(str1).Kind())
	fmt.Println("--- --- str1 --- ---")
	fmt.Println(str2)
	fmt.Println(reflect.TypeOf(str2))
	fmt.Println(reflect.ValueOf(str2).Kind())

	fmt.Println("--- --- e --- ---")
	xxx := e()
	fmt.Println(xxx)
	fmt.Println(reflect.TypeOf(xxx))
	fmt.Println(reflect.ValueOf(xxx).Kind())
}

func add (a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func e() (err error) {
	err = errors.New("e()")
	return
}