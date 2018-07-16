package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func initRedis(addr string, maxIdleConn, maxActiveConn int, idleTimeout time.Duration) (*redis.Pool) {
	pool = &redis.Pool{
		MaxIdle: maxIdleConn,
		MaxActive:maxActiveConn,
		IdleTimeout:idleTimeout,
		Dial:func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", addr)
		},
	}
	return pool
}