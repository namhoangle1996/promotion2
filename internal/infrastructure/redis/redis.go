package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"kk/config"
)

func NewRedisConn(redisConfig config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         redisConfig.RedisAddr,
		Password:     redisConfig.RedisPassword,
		DB:           redisConfig.DB,
		PoolSize:     redisConfig.PoolSize,
		PoolTimeout:  redisConfig.PoolTimeout,
		MinIdleConns: redisConfig.MinIdleConn,
	})

	if err := rdb.Ping(context.Background()); err != nil {
		panic(err)
	}

	return rdb
}
