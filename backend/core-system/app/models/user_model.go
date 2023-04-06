package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// User struct to describe User object.
type User struct {
	ID           uuid.UUID      `gorm:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"updated_at" json:"updated_at"`
	PhoneNumber  string         `gorm:"phone_number" json:"phone_number"`
	Email        string         `gorm:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string         `gorm:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserMetadata datatypes.JSON `gorm:"user_metadata" json:"user_metadata"`
	AppMetadata  datatypes.JSON `gorm:"app_metada" json:"app_metadata" `
}
