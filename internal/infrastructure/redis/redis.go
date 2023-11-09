package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewRedisConn() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Ping(context.Background()); err != nil {
		panic(err)
	}

	return rdb
}
