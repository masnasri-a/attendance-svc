package config

import (
	"fmt"
	"os"
	"time"

	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type RedisStore struct {
	Store            *persist.RedisStore
	DefaultCacheTime time.Duration
}

func GetRedisConfig(redisDb int) *redis.Client {
	godotenv.Load(".env")
	redisUri := os.Getenv("REDIS_URI")

	client := redis.NewClient(
		&redis.Options{
			Network:  "tcp",
			Addr:     redisUri,
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDb,
		},
	)
	return client
}

func GetRedisStore() *RedisStore {
	fmt.Println("Get Redis Store")
	client := GetRedisConfig(0)
	store := persist.NewRedisStore(client)
	return &RedisStore{
		Store:            store,
		DefaultCacheTime: time.Minute * 15,
	}
}
