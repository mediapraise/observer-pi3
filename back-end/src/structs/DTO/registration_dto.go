package DTO

import "time"

type RegistrationDTO struct {
	ID                 uint      `json:"id"`
	Board              string    `json:"board"`
	Vehicle            string    `json:"vehicle"`
	ExpirationDate     time.Time `json:"expiration_date"`
	RegistrationStatus string    `json:"registration_status"`
	Owner              string    `json:"owner"`
	CompanyID          string    `json:"company_id"`
}
