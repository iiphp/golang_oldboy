package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	path := "/Users/yangfan03/devspace/go/src/golang_oldboy/07day/03/example.txt"

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile fail")
	}
	fmt.Println(string(buf))
}
