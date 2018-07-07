package main

import (
	"fmt"
	"errors"
)

func initConf() (error) {
	return errors.New("initConf is Error")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovering")
		}
	}()

	err := initConf()
	if err != nil {
		panic(err)
	}
}
