package kafkaHandler

import (
	"gorm.io/gorm"
	"kk/config"
	"kk/internal/infrastructure/kafkaConn"
	"kk/internal/repositories/repositoriesImplement"
	"kk/internal/usecase/usecaseImple"
)

func NewKafkaConsumerHandler(kafkaConfig config.Kafka, db *gorm.DB) {
	repo := repositoriesImplement.NewPromotionRepo(db)
	promotionUsecase := usecaseImple.NewPromotionUsecase(repo)

	kafkaConsumer := kafkaConn.NewKafkaConn(kafkaConfig)

}
