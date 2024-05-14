package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID         uuid.UUID `gorm:"type:uuid;index"` // User who wrote the review
	DishID         uuid.UUID `gorm:"type:uuid;index"` // Dish being reviewed
	TenantID       uuid.UUID `gorm:"type:uuid;index"` // Tenant to which the dish belongs
	Content        string    `gorm:"not null"`        // The text content of the review
	Rating         float64   `gorm:"not null"`        // Numeric rating
	SentimentScore float64   `gorm:"type:float"`      // Sentiment score to be calculated
}

func (review *Review) BeforeCreate(tx *gorm.DB) (err error) {
	review.ID = uuid.New()
	return
}
