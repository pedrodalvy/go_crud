package database

import (
	"github.com/gofiber/fiber/v2/log"
	"go_crud/internal/products"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATA_SOURCE_NAME")), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect database")
	}

	err = db.AutoMigrate(&products.Product{})
	if err != nil {
		log.Error("failed to migrate database")
	}

	return db
}
