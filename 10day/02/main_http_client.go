package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	rsp, err := http.Get("http://www.baidu.com/")
	if nil != err {
		fmt.Println("http.Get is fail")
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if nil != err {
		fmt.Println("get body is fail")
		return
	}

	fmt.Println(string(body))
}
