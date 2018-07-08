package main

import (
	"fmt"
)

type Study interface {
	Domath()
}

type Work interface {
	Dowork()
}

type Student struct {
	Name string
	Age int
}

func (p *Student) Domath() {
	fmt.Println("do math")
}

// 判断一个对象，是否实现了某个接口
func main() {
	var std *Student
	var xyz interface{}
	xyz = std
	_, ok := xyz.(Study)
	fmt.Println(ok)
	if (ok) {
		fmt.Println("Study is succ")
	} else {
		fmt.Println("Study is fail")
	}

	// 注意，这里少了一个「冒号」，没有再用 :=
	// 因为，上面已经对 _ 和 ok 用 := 进行了初始化
	// 所以，这里直接赋值，不能再进行初始化了
	_, ok = xyz.(Work)
	if (ok) {
		fmt.Println("Work is succ")
	} else {
		fmt.Println("Work is fail")
	}
}
