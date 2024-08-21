package storage

import (
	"log"
	"os"

	"github.com/hazaloolu/blog-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	db.AutoMigrate(&model.User{}, &model.Post{})
	DB = db

}
