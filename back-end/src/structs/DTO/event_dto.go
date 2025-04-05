package DTO

import "time"

type EventDTO struct {
	ID        uint      `json:"id"`
	Date      time.Time `json:"date"`
	Board     string    `json:"board"`
	CompanyID string    `json:"company_id"`
}
