package config

import (
	"log"

	"backend-expense-app/internals/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Account{}, &models.Category{}, &models.SubCategory{})

	if err != nil {
		log.Fatal(err)
	}
}
