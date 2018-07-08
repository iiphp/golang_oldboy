package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {
	name string
	age int
}

type Policeman struct {
	Name string `json:"policeman_name"`
	Age int `json:"policeman_age"`
}

func main() {
	var std Student = Student{
		name: "Tom",
		age: 18,
	}

	// Student 的属性是小写的，所以 json 包访问不了 std 的 name 和 age 成员
	data, err := json.Marshal(std)
	if err != nil {
		fmt.Println("json.Marshal failed, err:", err)
	}
	fmt.Println(std)
	fmt.Println(string(data))

	var plc Policeman = Policeman{
		Name: "Jerry",
		Age: 33,
	}

	info, err := json.Marshal(plc)
	if err != nil {
		fmt.Println("json.Marshal failed, err:", err)
	}
	fmt.Println(plc)
	fmt.Println(string(info))
}
