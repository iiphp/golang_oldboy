package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b = [4]int{0, 1, 20, 300}
	testArray(b)
	testSlice()
	fmt.Println("--- --- b outer --- ---")
	fmt.Println(b)
	fmt.Println("--- --- b outer --- ---")
}

func testArray(b [4]int) {
	var a [4]int
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a).Kind())

	for i := 0; i < 3; i++ {
		a[i] = i * i * i;
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	b[3] = 10000
	fmt.Println("--- --- b inner --- ---")
	fmt.Println(b)
	fmt.Println("--- --- b inner --- ---")

	var c [2][3]int
	fmt.Println("--- --- c inner --- ---")
	fmt.Println(c)
	fmt.Println("--- --- c inner --- ---")
}

func testSlice() {
	fmt.Println("testSlice Bgn")
	var x []int
	var d = [5]int{0, 1, 20, 300, 4000}
	fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x).Kind())

	// slice 大小，由从数据的开始下标决定
	x = d[2:4]
	fmt.Println("x1 = ", x, len(x), cap(x))
	x = d[1:5]
	fmt.Println("x2 = ", x, len(x), cap(x))

	var y []int = make([]int, 10)
	fmt.Println("y1 = ", y, len(y), cap(y))


	z := "Hello World"
	var m = []byte(z)
	var n = []rune(z) // 能智能识别中文汉字
	m[1] = 'E'
	n[3] = 'L'
	fmt.Println("m1 =", string(m))
	fmt.Println("n1 =", string(n))
	fmt.Println("testSlice End")
}