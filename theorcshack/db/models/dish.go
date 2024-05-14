package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dish struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TenantID    uuid.UUID `gorm:"type:uuid;index"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       float64   `json:"price" binding:"required,gt=0"`
	ImageURL    string    `json:"image_url" binding:"required,url"`
}

func (dish *Dish) BeforeCreate(tx *gorm.DB) (err error) {
	dish.ID = uuid.New()
	return
}

var DB *gorm.DB
