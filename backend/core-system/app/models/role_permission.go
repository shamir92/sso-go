package models

import (
	"github.com/google/uuid"
)

// User struct to describe User object.
type RolePermission struct {
	APIPermissionID uuid.UUID     `gorm:"api_permission_id" json:"api_permission_id" validate:"required,uuid"`
	APIPermission   APIPermission `gorm: foreignKey:api_permission_id`
	RoleID          uuid.UUID     `gorm:"role_id" json:"role_id" validate:"required,uuid"`
	Role            Role          `gorm: foreignKey:role_id`
}
