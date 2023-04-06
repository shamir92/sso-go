package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type UserRole struct {
	gorm.Model
	UserID uuid.UUID `gorm:"user_id" json:"id" validate:"required,uuid"`
	User   User      `gorm:foreignKey:user_id`
	RoleID uuid.UUID `gorm:"role_id" json:"role_id"`
	Role   Role      `gorm:foreignKey:role_id`
}
