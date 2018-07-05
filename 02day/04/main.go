package main

import "fmt"

func main() {
	str := "hello world"
	fmt.Println(str)
	fmt.Println(str[2:])
	fmt.Println(str[2:8])
	fmt.Println(str[:8])

	str2 := reverse("1234567")
	fmt.Println(str2)
}

func reverse(s string) string {
	var rst []byte
	tmp := []byte(s)
	length := len(s)
	for i := 0; i < length; i++ {
		rst = append(rst, tmp[length-i-1])
	}
	return string(rst)
}