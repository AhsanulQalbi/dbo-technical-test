package config

import (
	"fmt"
	"os"

	"dbo-technical-test/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname =%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.Debug().AutoMigrate(
		models.User{},
		models.Customer{},
		models.Order{},
		models.Product{},
	)
	return db
}
