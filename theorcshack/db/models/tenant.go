package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name string    `gorm:"not null"`
}

func (tenant *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	tenant.ID = uuid.New()
	return
}
