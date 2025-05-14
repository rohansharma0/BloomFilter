package redisclient

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func InitRedis(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,     // e.g. "localhost:6379"
		Password: password, // "" if no password
		DB:       db,       // 0 is default DB
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Redis connected.")
}
