package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	// 注意，这里是单引号的 \n ，是 byte
	str, err := r.ReadString('\n')
	if (nil != err) {
		fmt.Println("ReadString fail")
	}
	fmt.Println(str)
}
