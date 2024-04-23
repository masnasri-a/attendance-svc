package utils

import (
	"attendance-svc/src/config"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCacheData(c *gin.Context, value interface{}, expiration time.Duration, useAuth bool) {
	fmt.Println("Set Cache Data")
	marshallValue, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error Marshal Data: ", err)
	}
	configRedis := config.GetRedisConfig(0)
	key := c.Request.RequestURI
	if !useAuth {
		configRedis.SetEX(context.Background(), key, marshallValue, expiration)
		return
	}
	authHeader := c.GetHeader("Authorization")
	userId, _ := VerifyToken(authHeader)
	key = key + "?userId=" + userId
	fmt.Println("key: ", key)
	fmt.Println("value: ", value)
	if err := configRedis.SetEX(context.Background(), key, marshallValue, expiration).Err(); err != nil {
		fmt.Println("Error Set Cache Data: ", err)
	}
	fmt.Println("Set Cache Data Done")
}

func GetCacheData(c *gin.Context, useAuth bool) (interface{}, bool) {
	fmt.Println("Get Cache Data")
	configRedis := config.GetRedisConfig(0)
	key := c.Request.RequestURI
	if !useAuth {
		value, err := configRedis.Get(context.Background(), key).Result()
		if err != nil {
			return nil, false
		}
		return value, true
	}
	authHeader := c.GetHeader("Authorization")
	userId, _ := VerifyToken(authHeader)
	key = key + "?userId=" + userId
	value, err := configRedis.Get(context.Background(), key).Result()
	fmt.Println("key: ", key)
	if err != nil {
		return nil, false
	}
	unmarshallValue := make(map[string]interface{})
	json.Unmarshal([]byte(value), &unmarshallValue)
	return unmarshallValue, true
}
