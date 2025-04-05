package DTO

type ClientDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Document  string `json:"document"`
	CompanyID uint   `json:"company_id"`
}
