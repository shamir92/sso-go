package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type Role struct {
	ID          uuid.UUID `gorm:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at" json:"updated_at"`
	Name        string    `gorm:"name" json:"name" validate:"required,lte=255"`
	Description string    `gorm:"description" json:"description"`
}
