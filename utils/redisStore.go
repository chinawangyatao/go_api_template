package utils

import (
	"context"
	"go_admin/pkg/redis"
	"time"
)

var (
	ctx       = context.Background()
	loginCode = "redisLoginCode"
)

type RedisStore struct {
}

// Set 存验证码
func (s RedisStore) Set(id string, value string) error {
	key := loginCode + id
	err := redis.DbRedis.Set(ctx, key, value, time.Minute*10).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get 取验证码
func (s RedisStore) Get(id string, clear bool) string {
	key := loginCode + id
	val, err := redis.DbRedis.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

// Verify 校验验证码
func (s RedisStore) Verify(id, captcha string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == captcha
}
