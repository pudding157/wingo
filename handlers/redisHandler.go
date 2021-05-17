package handlers //เปรียบเสมือน namespace c#

import (
	"log"

	// "net/http"
	"github.com/go-redis/redis"
)

type RedisHandler struct {
	DB *redis.Client
}

func (h *RedisHandler) Initialize() {
	log.Printf("hello redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	log.Printf(pong, err)
	h.DB = client

}
