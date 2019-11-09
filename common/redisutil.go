package common

import (
	"github.com/go-redis/redis/v7"
	"log"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})

	pong, err := RedisClient.Ping().Result()
	log.Println(pong, err)
}
