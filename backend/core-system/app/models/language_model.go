package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type Language struct {
	ID           uuid.UUID      `gorm:"id,primaryKey:type:uuid;autoIncrement:false" json:"id"`
	CreatedAt    time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"deleted_at" json:"deleted_at"`
	Name         string         `gorm:"name" json:"name"`
	DisplayName  string         `gorm:"display_name" json:"display_name" validate:"required"`
	LanguageCode string         `gorm:"column:language_code;unique" json:"language_code" validate:"required"`
}

type LanguageRequest struct {
	Name         string `json:"name" validate:"required"`
	DisplayName  string `json:"display_name" validate:"required"`
	LanguageCode string `json:"language_code" validate:"required"`
}

type LanguageSetJson []struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
