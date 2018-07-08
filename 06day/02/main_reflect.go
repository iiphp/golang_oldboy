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

func main() {
	std := Student{Name:"Tom", Age:21}
	foo(std)
}
