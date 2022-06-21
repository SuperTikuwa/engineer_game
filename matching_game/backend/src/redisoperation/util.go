package redisoperation

import (
	"fmt"
	"hash/fnv"
	"log"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func RedisConnect() redis.Conn {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	return conn
}

func GenerateGameID() string {
	conn := RedisConnect()
	defer conn.Close()

	h := fnv.New32()
	h.Write([]byte(time.Now().String()))
	// h.Write([]byte("hoge"))
	sum := strconv.Itoa(int(h.Sum32()))

	len, err := redis.Int(conn.Do("LLEN", sum))
	if err != nil {
		log.Println(err)
	}

	if len != 0 {
		// log.Fatal("Game already exists.")
		sum = GenerateGameID()
	}

	return sum
}

func JoinGame(gameID, playerID string) {
	conn := RedisConnect()
	defer conn.Close()

	fmt.Println(gameID, playerID)

	_, err := conn.Do("LPUSH", gameID, playerID)
	if err != nil {
		log.Println(err)
	}
}
