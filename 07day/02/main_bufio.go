package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	f, err := os.Open("/Users/yangfan03/devspace/go/src/golang_oldboy/07day/02/example.txt")
	if nil != err {
		fmt.Println("Open file fail")
		return
	}
	// 注意，这里要记得随时关闭文件
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, e := r.ReadString('\n')

		runeData := []rune(s)
		for _, v := range runeData {
			fmt.Println(string(v))
		}

		if io.EOF == e {
			break
		}
	}
}
