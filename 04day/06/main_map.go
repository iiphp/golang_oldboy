package main

import "fmt"

func main()  {
	testMap()
	testMap2()
}

func testMap()  {
	fmt.Println()
	fmt.Println("--- --- testMap bgn --- ---")

	a := make(map[string]int, 10) // 如果超过了 10 ，会再分配，跟 slice 比较类似
	a["a"] = 1
	a["b"] = 2
	fmt.Println("a1 = ", a)

	val, ok := a["a"]
	if ok {
		fmt.Println("a is OK", val)
	}

	for k, v := range a {
		fmt.Println("k is:", k, " v is:", v)
	}

	delete(a, "a")

	for k, v := range a {
		fmt.Println("k is:", k, " v is:", v)
	}
	fmt.Println("--- --- testMap end --- ---")
	fmt.Println()
}

func testMap2()  {
	fmt.Println()
	fmt.Println("--- --- testMap2 bgn --- ---")

	b := make(map[int]map[string]string, 10)
	b[0] = make(map[string]string)
	b[0]["abc"] = "ABC"
	b[1] = make(map[string]string)
	b[1]["def"] = "DEF"
	fmt.Println("b = ", b)

	fmt.Println("--- --- testMap2 end --- ---")
	fmt.Println()
}