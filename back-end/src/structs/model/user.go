package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name               string    `gorm:"type:varchar(100);not null"`
	Email              string    `gorm:"type:varchar(100);not null;unique"`
	Password           string    `gorm:"not null"`
	Verified           bool      `gorm:"not null"`
	VerificationCode   string    `gorm:"type:varchar(50)"`
	ExpireVerification time.Time `gorm:"not null"`
	CompanyID          uint      `gorm:"not null"` // Foreign key
	Company            Company   `gorm:"foreignKey:CompanyID"`
}
