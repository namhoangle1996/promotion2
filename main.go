package main

import (
	"kk/config"
	"kk/internal/handler/httpHandler"
	"kk/internal/infrastructure/postgresql"
	"kk/internal/infrastructure/redis"
)

func main() {
	conf := config.NewConfig()
	_ = redis.NewRedisConn(conf.Redis)

	db := postgresql.NewConn(conf.Postgresql)

	//http handler
	httpHandler.NewHttpServer(conf.HTTP, db)

	//kafka consumer handler

	select {}
}
