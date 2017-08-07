package redis

import (
	"log"
	"lotteryGame/config"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword, // no password set
		DB:       config.RedisDB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("connect redis error :", err)
	}
	log.Println(pong, err)
}
