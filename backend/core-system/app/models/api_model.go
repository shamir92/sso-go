package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type API struct {
	ID               uuid.UUID      `gorm:"id,primaryKey:type:uuid;autoIncrement:false" json:"id"`
	CreatedAt        time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"deleted_at" json:"deleted_at"`
	Identifier       string         `gorm:"identifier" json:"identifier" validate:"required"`
	SigningAlgorithm string         `gorm:"signing_algorithm" json:"signing_algorithm" validate:"required"`
}

type APIRequest struct {
	Identifier       string `json:"identifier" validate:"required"`
	SigningAlgorithm string `json:"signing_algorithm" validate:"required"`
}

func IsValidSigningAlgorithm(ident string) bool {
	switch ident {
	case
		"HS256", "RS256":
		return true
	}
	return false
}
