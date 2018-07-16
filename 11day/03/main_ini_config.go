package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
)

func main() {
	iniFileName := "/usr/local/etc/php/7.1/php.ini"

	iniConfig, err := config.NewConfig("ini", iniFileName)
	if nil != err {
		fmt.Println("config.NewConfig failed", err)
		return
	}

	shortOpenTag, err := iniConfig.Bool("PHP::short_open_tag")
	if nil != err {
		fmt.Println("iniConfig get short_open_tag failed")
		return
	}
	fmt.Println("short_open_tag:", shortOpenTag)


	engine, err := iniConfig.Bool("PHP::engine")
	if nil != err {
		fmt.Println("iniConfig get engine failed")
		return
	}
	fmt.Println("engine:", engine)
}
