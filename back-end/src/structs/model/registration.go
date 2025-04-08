package model

import (
	"time"

	"gorm.io/gorm"
)

type Registration struct {
	// Embeds gorm.Model to include fields like ID, CreatedAt, UpdatedAt, and DeletedAt
	gorm.Model
	Board   string `gorm:"type:varchar(100);not null"`
	Vehicle string `gorm:"type:varchar(100);not null"`
	ExpirationDate time.Time `gorm:"type:date;not null"`
	RegistrationStatus string `gorm:"type:varchar(100);not null"`
	Owner   string `gorm:"type:varchar(100);not null"`
	CompanyID string `gorm:"type:varchar(100);not null"`
}

