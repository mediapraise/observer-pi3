package model

import (
	"time"

	"gorm.io/gorm"
)

type HistoryPayment struct {
	gorm.Model
	Board		  string `gorm:"type:varchar(100);not null"`
	PaymentDate time.Time `gorm:"type:date;not null"`
	PaymentStatus string `gorm:"type:varchar(100);not null"`
	PaymentAmount int64 `gorm:"type:bigint;not null"`
	CompanyID uint `gorm:"not null"`
}