package main

import "fmt"

type Student struct {
	Name string
	Age int
	Next *Student
}


func main() {
	std1 := Student{Name:"Tom", Age:18, Next:nil}
	std2 := Student{Name:"Peter", Age:21, Next:&std1}
	std3 := Student{Name:"Jerry", Age:39, Next:&std2}

	stdx := std3
	for {
		fmt.Println(stdx)
		if nil == stdx.Next {
			break
		}
		stdx = *(stdx.Next)
	}

	stdz := &std3
	for {
		fmt.Println(*stdz)
		if (stdz.Next == nil) {
			break
		}
		stdz = stdz.Next
	}
}
