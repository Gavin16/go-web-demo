package redisgo

import (
	redis "github.com/gomodule/redigo/redis"
	"time"
)

var redisPool *redis.Pool

var redisServer = "127.0.0.1:6379"

func init() {
	redisPool = &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 6 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisServer)
		},
	}
}

// GetConn 获取redis链接
func GetConn() redis.Conn {
	return redisPool.Get()
}
