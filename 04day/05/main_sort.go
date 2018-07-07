package main

import (
	"sort"
	"fmt"
)

func main() {
	testSortInt()
	testSortStr()
}

func testSortInt() {
	fmt.Println()
	fmt.Println("--- --- testSortInt bgn --- ---")
	a := [...]int{3, 9172, 671, 44}
	fmt.Println("a init is ", a)

	// 这里要传给 sort 切片进去
	// 因为如果给 sort.Ints 函数传数组，那么是「值传递」，在里面干完活后，并不会影响外面的数组 a
	// 所以要传给 sort.Ints 函数一个指向数组 a 的指针，那么就是个切片

	b := a[:]
	sort.Ints(b)
	fmt.Println("a is sorted ", a)
	fmt.Println("--- --- testSortInt end --- ---")
	fmt.Println()
}

func testSortStr()  {
	fmt.Println()
	fmt.Println("--- --- testSortInt bgn --- ---")
	// sort.Strings()
	fmt.Println("--- --- testSortInt end --- ---")
	fmt.Println()
}