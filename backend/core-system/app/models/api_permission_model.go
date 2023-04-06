package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type APIPermission struct {
	ID          uuid.UUID      `gorm:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"updated_at" json:"updated_at"`
	Scope       string         `gorm:"scope" json:"scope"`
	Description string         `gorm:"description" json:"description"`
	APIID       uuid.UUID      `gorm:"api_id" json:"api_id" validate:"required,uuid"`
	API         API            `gorm: foreignKey:api_id`
	DeletedAt   gorm.DeletedAt `gorm:"deleted_at" json:"deleted_at"`
}
