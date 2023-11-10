package kafkaConn

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kk/config"
)

func NewKafkaConn(kafkaConfig config.Kafka) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaConfig.BoostrapServers,
		"group.id":          kafkaConfig.GroupId,
	})

	if err != nil {
		panic(err)
	}
	return c
}
