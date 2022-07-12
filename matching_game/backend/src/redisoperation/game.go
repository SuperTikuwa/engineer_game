package redisoperation

import (
	"hash/fnv"
	"log"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

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

func CreateGame(gameID, player string) {
	conn := RedisConnect()
	defer conn.Close()
	_, err := conn.Do("SET", "games:"+gameID+":player1", player)
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Do("SET", "games:"+gameID+":player2", "")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Do("SET", "games:"+gameID+":isEmpty", true)
	if err != nil {
		log.Println(err)
	}
}

func FindEmptyGame() string {
	conn := RedisConnect()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "games:*:isEmpty"))
	if err != nil {
		log.Fatal(err)
	}

	for _, key := range keys {
		val, err := redis.Bool(conn.Do("GET", key))
		if err != nil {
			log.Fatal(err)
		}
		if val {
			return key
		}
	}

	return ""
}

func JoinGame(gameID, playerID string) {
	conn := RedisConnect()
	defer conn.Close()

	_, err := conn.Do("SET", "games:"+gameID+":player2", playerID)
	if err != nil {
		log.Println(err)
	}

	_, err = conn.Do("SET", "games:"+gameID+":isEmpty", false)
	if err != nil {
		log.Println(err)
	}
}

func GetGames() []string {
	conn := RedisConnect()
	games, err := redis.Strings(conn.Do("KEYS", "*"))

	if err != nil {
		log.Fatal(err)
	}

	return games
}
