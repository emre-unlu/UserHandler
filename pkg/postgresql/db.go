package postgresql

import (
	"fmt"
	"github.com/emre-unlu/GinTest/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=newuser password=newpassword dbname=newdatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
		os.Exit(1)
	}
	db.AutoMigrate(&models.User{})

	DB = db
	fmt.Println("Connected to database ")
}
