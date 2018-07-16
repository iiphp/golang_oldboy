package main

import (
	"context"
	"fmt"
)

// 通用的参数用 ctx 传递，比如 req_id 等

func main() {
	ctx := context.WithValue(context.Background(), "req_id", "a1b2c3d4")

	doSth(ctx)
}

func doSth(ctx context.Context) {
	req_id, ok := ctx.Value("req_id").(string)
	if !ok {
		fmt.Println("ctx.Value(req_id) failed")
		return
	}

	fmt.Printf("req_id: %s", req_id)
}