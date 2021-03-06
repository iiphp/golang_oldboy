package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string
	Age int
}


// 注意 TypeOf 和 ValueOf.kind() 的区别
// 三者相互转化，「变量」和 「interface()」 相互转，「Reflect.Value」和「interface()」相互转
func foo(input interface{}) {
	x := reflect.TypeOf(input)
	fmt.Println(x)

	v := reflect.ValueOf(input)
	k := v.Kind()
	fmt.Println(v, k)

	iv := v.Interface()
	std, ok := iv.(Student)
	if (ok) {
		fmt.Println(std)
	} else {
		fmt.Println("fail")
	}
}

// 反射方法，获取变量的值，转成普通变量
func val(input interface{}) {
	v := reflect.ValueOf(input)
	fmt.Println(v, reflect.TypeOf(v))
	x := v.Int()
	fmt.Println(x, reflect.TypeOf(x))
}

func main() {
	std := Student{Name:"Tom", Age:21}
	foo(std)

	i := 10
	val(i)
}
