package config

import (
	"errors"
	"fmt"
	"os"

	"dbo-technical-test/models"

	log "github.com/sirupsen/logrus"
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
	if err == nil && db.Migrator().HasTable(&models.User{}) {
		if err := db.Where("email", "superadmin@dbo.com").First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			user := []*models.User{
				{Fullname: "superadmin", Password: "$2a$08$0WveZ9.JfFGSnf9H5SBgfev8gJ3SbeqDhLexLQkks/WmYnPxTZLnS", Email: "superadmin@dbo.com",
					Role: "Super Admin", Address: "Tangerang"},
			}
			if err := db.Create(&user).Error; err != nil {
				log.Errorf("[seed super admin] err: %s", err)
			}
		}
	}
	return db
}
