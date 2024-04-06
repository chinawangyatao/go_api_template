// Package redis 初始化 redis
package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go_admin/common/config"
)

var (
	DbRedis *redis.Client
)

func RedisInit() error {
	config.Init()
	DbRedis = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
	_, err := DbRedis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("ping err %v", err)
		panic(err)
	}
	return nil
}
