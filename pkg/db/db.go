package db

import (
	"log"

	"github.com/bishal-dd/go-server/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := "host=localhost user=bishal password=bishal dbname=godatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&model.User{})

    return db
}