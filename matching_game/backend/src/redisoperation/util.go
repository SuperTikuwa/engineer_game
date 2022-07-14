package redisoperation

import (
	"github.com/gomodule/redigo/redis"
)

func RedisConnect() redis.Conn {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	return conn
}
