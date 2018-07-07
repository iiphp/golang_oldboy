package main

import (
	"time"
	"fmt"
	"sync"
)

// 定义一个全局的锁
var lockMutex sync.Mutex
var lockRWMutex sync.RWMutex

func main()  {
	testMutex()
	testRWMutex()
}


func testRWMutex()  {
	bgn_time := time.Now()

	// 设置 2 个 goroutine 负责写
	for i := 0; i < 2; i++ {
		wkey := fmt.Sprintf("w-rw%d", i)
		rkey := fmt.Sprintf("r-rw%d", i)
		go writeRWMutex(wkey, bgn_time)
		go readRWMutex(rkey, bgn_time)
	}

	time.Sleep(time.Second * 12)
	// 设置 5 个 goroutine 负责读
}

func writeRWMutex(key string, bgn_time time.Time)  {
	for i := 0; i < 3; i++ {
		lockRWMutex.Lock()
		time.Sleep(time.Millisecond * 1000)
		elapsed := time.Since(bgn_time)
		fmt.Println("[", key, "] write [", i, "]is done", elapsed)
		lockRWMutex.Unlock()
	}
}

func readRWMutex(key string, bgn_time time.Time)  {
	for i := 0; i < 10; i++ {
		lockRWMutex.RLock()
		time.Sleep(time.Millisecond * 200)
		elapsed := time.Since(bgn_time)
		fmt.Println("[", key, "] read [", i, "]is done", elapsed)
		lockRWMutex.RUnlock()
	}
}

func testMutex()  {
	bgn_time := time.Now()

	// 设置 2 个 goroutine 负责写
	for i := 0; i < 2; i++ {
		wkey := fmt.Sprintf("w%d", i)
		rkey := fmt.Sprintf("r%d", i)
		go writeMutex(wkey, bgn_time)
		go readMutex(rkey, bgn_time)
	}

	time.Sleep(time.Second * 12)
	// 设置 5 个 goroutine 负责读
}

func writeMutex(key string, bgn_time time.Time)  {
	for i := 0; i < 3; i++ {
		lockMutex.Lock()
		time.Sleep(time.Millisecond * 1000)
		elapsed := time.Since(bgn_time)
		fmt.Println("[", key, "] write [", i, "]is done", elapsed)
		lockMutex.Unlock()
	}
}

func readMutex(key string, bgn_time time.Time)  {
	for i := 0; i < 10; i++ {
		lockMutex.Lock()
		time.Sleep(time.Millisecond * 200)
		elapsed := time.Since(bgn_time)
		fmt.Println("[", key, "] read [", i, "]is done", elapsed)
		lockMutex.Unlock()
	}
}