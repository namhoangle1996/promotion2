package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kk/config"
	"kk/internal/model"
	"log"
)

func NewConn(postgresql config.Postgresql) *gorm.DB {
	host := postgresql.PostgresqlHost
	user := postgresql.PostgresqlUser
	pw := postgresql.PostgresqlPassword
	dbName := postgresql.PostgresqlDBName
	port := postgresql.PostgresqlPort
	sslMode := postgresql.PostgresqlSSLMode

	dbURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Ho_Chi_Minh", host, user, pw, dbName, port, sslMode)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{DryRun: true, PrepareStmt: true})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	err = db.AutoMigrate(&model.Promotion{}, &model.PromotionUsage{})
	if err != nil {
		panic(err)
	}
	return db
}
