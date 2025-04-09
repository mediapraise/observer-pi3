package DTO

type CompanyDTO struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	OwnerID uint   `json:"owner_id"`
}
