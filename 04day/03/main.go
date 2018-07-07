package main

import "fmt"

func main() {
	// 一个指向切片的指针，指向的切片是空的。
	s1 := new([]int)
	fmt.Println("--- --- No.1 --- ---")
	fmt.Println(s1)
	fmt.Println(*s1) // 这个指针指向的数据，是一个空的切片
	fmt.Println(&s1) // 这个指针变量的地址

	// 初始化一个切片
	s2 := make([]int, 3)
	fmt.Println("--- --- No.2 --- ---")
	fmt.Println(s2)
	// 此行会报错 fmt.Println(*s2)
	fmt.Println(&s2)

	// 把刚才那个指向切片的指针 s1 进行赋值，让它指向一个新的切片
	*s1 = make([]int, 5)
	fmt.Println("--- --- No.3 --- ---")
	fmt.Println(s1)
	fmt.Println(*s1)
	fmt.Println(&s1)
}