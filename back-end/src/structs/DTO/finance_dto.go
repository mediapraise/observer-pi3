package DTO

import "time"

// HistoryPaymentDTO represents the DTO for HistoryPaymant model
type HistoryPaymentDTO struct {
	ID            uint      `json:"id"`
	Board         string    `json:"board"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentStatus string    `json:"payment_status"`
	PaymentAmount int64     `json:"payment_amount"`
	CompanyID     string      `json:"company_id"`
}
