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

type EventDTO struct {
	ID        uint      `json:"id"`
	Date      time.Time `json:"date"`
	Board     string    `json:"board"`
	CompanyID string    `json:"company_id"`
	Event     bool      `json:"event"`
}
