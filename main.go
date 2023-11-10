package main

import (
	"kk/config"
	"kk/internal/infrastructure/redis"
)

func main() {
	conf := config.NewConfig()

	_ = redis.NewRedisConn(conf.Redis)

	select {}
}
