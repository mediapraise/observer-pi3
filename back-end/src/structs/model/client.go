package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);not null;unique"`
	Phone     string `gorm:"type:varchar(100);not null;unique"`
	Document  string `gorm:"type:varchar(100);not null;unique"`
	CompanyID uint   `gorm:"not null"` // Foreign key for Company
}
