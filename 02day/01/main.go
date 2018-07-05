package main

import "fmt"

func main ()  {
	list(5)
}

func list(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d + %d = %d\n", i, n, n + i)
	}
}

