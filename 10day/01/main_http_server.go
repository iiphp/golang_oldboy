package main

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle hello")
	fmt.Fprint(w, "hello world")
}

func main()  {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe("0.0.0.0:8937", nil) // 可以访问 http://localhost:8937/hello
	if nil != err {
		fmt.Println("ListenAndServe Fail")
	}
}