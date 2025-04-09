package DTO

type UserDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
	IsAdmin   bool   `json:"is_admin"`
	CompanyID uint   `json:"company_id"`
}
