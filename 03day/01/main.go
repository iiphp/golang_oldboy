package main

import (
	"fmt"
	"strings"
)

func doUrl(url string) string {
	if (strings.HasPrefix(url, "http:")) {
		return url
	}
	return fmt.Sprintf("http:%s", url)
}

func main() {
	var url string
	fmt.Scanf("%s", &url)
	fmt.Println(url)

	url = doUrl(url)
	fmt.Println(url)
}
