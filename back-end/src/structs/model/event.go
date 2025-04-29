package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Date      time.Time `gorm:"type:date;not null"`
	Board     string    `gorm:"type:varchar(100);not null"`
	CompanyID string    `gorm:"type:varchar(100);not null"`
	Event     bool      `gorm:"type:bool;not null"`
}
