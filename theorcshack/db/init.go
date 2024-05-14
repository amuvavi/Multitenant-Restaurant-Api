package db

import (
	"fmt"
	"log"
	"os"
	"theorcshack/db/migrations"
	"theorcshack/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	models.DB = db

	migrations.Migrate(db)
}
