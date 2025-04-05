package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name    string   `gorm:"column:name;type:varchar(100)"`
	Address string   `gorm:"column:address;type:varchar(255)"`
	Phone   string   `gorm:"column:phone;type:varchar(20)"`
	Email   string   `gorm:"column:email;type:varchar(100)"`
	Users   []User   `gorm:"foreignKey:CompanyID"`
	Clients []Client `gorm:"foreignKey:CompanyID"`
	OwnerID uint     `gorm:"column:owner_id"`
	Owner   User     `gorm:"foreignKey:OwnerID"`
}
