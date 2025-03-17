package model

import (
	"time"

	"gorm.io/gorm"
)

type Boards struct {
	gorm.Model
	Code      string    `gorm:"type:varchar(100);not null"`
	Vehicle   string    `gorm:"type:varchar(100);not null"`
	Owner     string    `gorm:"type:varchar(100);not null"`
	EntryDate time.Time `gorm:"not null"`
	ExitDate  time.Time `gorm:"not null"`
}
