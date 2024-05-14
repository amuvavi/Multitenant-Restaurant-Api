package migrations

import (
	"theorcshack/db/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"") // Enable UUID extension for Postgres
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Dish{})
	db.AutoMigrate(&models.Tenant{})
	db.AutoMigrate(&models.Review{})

	db.Exec("CREATE INDEX IF NOT EXISTS idx_dishes_tenant_id ON dishes (tenant_id)")

	// Add tenant_id column to users table
	AddTenantIDToUsers(db)

}
