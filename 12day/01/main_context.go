package main

import (
	"context"
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
)

type Rsp struct {
	r   *http.Response
	err error
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1)
	defer cancel()

	tr := &http.Transport{}
	client := &http.Client{Transport:tr}
	c := make(chan Rsp, 1)
	req, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	if nil != err {
		fmt.Println("http.NewRequest faild")
	}

	go func() {
		resp, err := client.Do(req)
		if nil != err {
			fmt.Println("client Do failed", err)
		}
		pack := Rsp{r:resp, err:err}
		c <- pack
	}()

	select {
	case <- ctx.Done():
		tr.CancelRequest(req) // 已经弃用，因为取消不了 HTTP/2 的请求 https://www.bookstack.cn/read/gctt-godoc/Translate-net-http-http.md#Transport.CancelRequest
		<- c
		fmt.Println("Timeout")
	case res := <- c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Response: %s", out)
	}
}
