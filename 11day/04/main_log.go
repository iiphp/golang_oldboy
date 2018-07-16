package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func main()  {
	logConfig := make(map[string]interface{})
	logConfig["filename"] = "/tmp/tmp.log"
	logConfig["level"]    = logs.LevelDebug

	enJsonLogConfig, err := json.Marshal(logConfig)
	if nil != err {
		fmt.Println("json.Marshal logConfig failed", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(enJsonLogConfig))

	logs.Debug("this is a debug message")
	logs.Info("this is a info message")
	logs.Error("this is a error message")
}
