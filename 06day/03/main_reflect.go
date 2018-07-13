package main

import (
	"reflect"
	"fmt"
)

// 注意 Elem() 方法，获得反射指针
func foo(input interface{}) {
	v := reflect.ValueOf(input)
	v.Elem().SetInt(200)

	r := v.Elem().Int()
	fmt.Println(r, reflect.TypeOf(r))
}

// 反射修改变量的值
func main()  {
	i := 10

	fmt.Println("i1 =", i)
	// 注意：这里要传「地址」，因为我们要修改变量的值
	foo(&i)
	fmt.Println("i2 =", i)
}
