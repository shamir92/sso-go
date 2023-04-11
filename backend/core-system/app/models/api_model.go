package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type API struct {
	ID                         uuid.UUID      `gorm:"id,primaryKey:type:uuid;autoIncrement:false" json:"id"`
	CreatedAt                  time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt                  time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt                  gorm.DeletedAt `gorm:"deleted_at" json:"deleted_at"`
	Name                       string         `gorm:"name" json:"name"`
	Identifier                 string         `gorm:"identifier" json:"identifier" validate:"required"`
	SigningAlgorithm           string         `gorm:"signing_algorithm" json:"signing_algorithm" validate:"required"`
	TokenExpiration            int            `gorm:"token_expiration"  json:"token_expiration"`
	TokenExpirationForBrowser  int            `gorm:"token_expiration_for_browser" json:"token_expiration_for_browser"`
	EnableRBAC                 bool           `gorm:"enable_rbac" json:"enable_rbac"`
	AddPermissionToAccessToken bool           `gorm:"add_permission_to_access_token" json:"add_permission_to_access_token"`
	AllowSkippingUserConsent   bool           `gorm:"allow_skipping_user_consent" json:"allow_skipping_user_consent"`
	AllowOfflineAccess         bool           `gorm:"allow_offline_access" json:"allow_offline_access"`
}

type APIRequest struct {
	Identifier       string `json:"identifier" validate:"required"`
	SigningAlgorithm string `json:"signing_algorithm" validate:"required"`
}

type APISettingRequest struct {
	Name                       string `validate:"required,alphanum" json:"name"`
	Identifier                 string `validate:"required" json:"identifier"`
	TokenExpiration            int    `validate:"numeric,gte=60" json:"token_expiration"`
	TokenExpirationForBrowser  int    `validate:"numeric,gte=60" json:"token_expiration_for_browser"`
	EnableRBAC                 string `validate:"required,boolean,oneof=false true" json:"enable_rbac"`
	AddPermissionToAccessToken string `validate:"required,boolean,oneof=false true" json:"add_permission_to_access_token"`
	AllowSkippingUserConsent   string `validate:"required,boolean,oneof=false true" json:"allow_skipping_user_consent"`
	AllowOfflineAccess         string `validate:"required,boolean,oneof=false true" json:"allow_offline_access"`
}

func IsValidSigningAlgorithm(ident string) bool {
	switch ident {
	case
		"HS256", "RS256":
		return true
	}
	return false
}
