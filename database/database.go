package database

import (
	"fmt"
	"goworktoday/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// postgres initialization and connections
func InitDB() *gorm.DB {
	dsn := priv
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	// Migrate models and error handling
	if err := db.AutoMigrate(
		&models.User{},
		&models.MainAccount{},
		&models.SavingsAccount{},
		&models.CardAccount{},
		&models.Blacklist{},
	); err != nil {
		log.Fatalf("Failed to perform database migrations: %v", err)
	}

	DB = db

	return db
}
