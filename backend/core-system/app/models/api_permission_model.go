package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type APIPermission struct {
	ID          uuid.UUID `gorm:"id" json:"id" validate:"omitempty,uuid"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at" `
	UpdatedAt   time.Time `gorm:"updated_at" json:"updated_at"`
	Scope       string    `gorm:"scope" json:"scope" validate:"required,alphanum"`
	Description string    `gorm:"description" json:"description" validate:"required,alphanum"`
	APIID       uuid.UUID `gorm:"api_id" json:"api_id" validate:"omitempty,uuid"`
	// API         API
	DeletedAt gorm.DeletedAt `gorm:"deleted_at" json:"deleted_at"`
}

type APIPermissionRequest struct {
	Scope       string `json:"scope" validate:"required"`
	Description string `json:"description" validate:"required,alphanum"`
}
