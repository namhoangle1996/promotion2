package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	AppVersion string
	HTTP       HTTP
	KafKa      Kafka
	Redis      Redis
	Postgresql Postgresql
}

type Kafka struct {
	BoostrapServers string
	GroupId         string
}

type Redis struct {
	RedisAddr      string
	RedisPassword  string
	RedisDefaultDB int
	MinIdleConn    int
	PoolSize       int
	PoolTimeout    time.Duration
	DB             int
}

type HTTP struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

type Postgresql struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDBName   string
	PostgresqlSSLMode  string
}

func NewConfig() (conf Config) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})

	err = viper.Unmarshal(&conf)

	return conf
}
