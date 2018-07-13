package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name string `json:"name_name"`
	Age int
	Score int
}

func main() {
	encode()
	decode()
}

func encode() {
	std := Student{Name:"Tom", Age:18, Score:127}
	fmt.Println(std)

	rsp, err := json.Marshal(std)
	if nil != err {
		fmt.Println("json.Marshal fail")
	}

	fmt.Println(string(rsp))
}

func decode()  {
	var std Student
	str := "{\"name_name\":\"Tom\",\"Age\":18,\"Score\":127}"
	// 注意，反序列化，接收的参数要传「引用」
	json.Unmarshal([]byte(str), &std)
	fmt.Println(std)
}