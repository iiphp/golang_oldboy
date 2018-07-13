package main

import (
	"fmt"
	"flag"
)

func main()  {
	var size int
	var path string

	flag.IntVar(&size, "s", 10, "pls input size")
	flag.StringVar(&path, "p", "", "pls input path")
	flag.Parse()

	fmt.Println("size = ", size)
	fmt.Println("path = ", path)
}
