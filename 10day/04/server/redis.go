package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initRedis(addr string, maxIdleConn, maxActiveConn int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle: maxIdleConn,
		MaxActive:maxActiveConn,
		IdleTimeout:idleTimeout,
		Dial:func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", addr)
		},
	}
}

func GetConn() redis.Conn {
	return pool.Get()
}

func PutConn(conn redis.Conn) {
	conn.Close()
}