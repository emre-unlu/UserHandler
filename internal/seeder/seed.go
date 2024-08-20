package seeder

import (
	"fmt"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/internal/utils"
	"gorm.io/gorm"
	"log"
)

func SeedUser(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count == 0 {
		password, err := utils.HashPassword("admin123")
		if err != nil {
			panic(err)
		}

		user := models.User{
			Name:     "Admin",
			Surname:  "User",
			Email:    "admin@example.com",
			Phone:    "+391234567890",
			Password: password,
			Status:   models.StatusActive}

		result := db.Create(&user)
		if result.Error != nil {
			log.Fatalf("Failed to create initial user: %v", result.Error)
		}

		fmt.Printf("Initial user created with email: admin@example.com and password: %s\n", "admin123")
	} else {
		fmt.Println("Users already exist, skipping seeding.")
	}
}
