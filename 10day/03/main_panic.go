package main

import (
	"net/http"
	"fmt"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "hello, world")
}

func logPainc(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); nil != e {
				log.Printf("logPainc")
			}
		}()
		handle(writer, r)
	}

}

func main() {
	http.HandleFunc("/hello", logPainc(hello))

	if err := http.ListenAndServe(":8643", nil); nil != err {
		fmt.Println("ListenAndServe is fail")
	}
}
