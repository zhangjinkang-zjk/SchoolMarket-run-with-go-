package database

import (
	"context"
	"fmt"

	"SchoolMarket-run-with-go-/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.RD_HOST, cfg.RD_PORT)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis连接失败: %v", err))
	}

	fmt.Println("Redis连接成功")
	return RedisClient
}
