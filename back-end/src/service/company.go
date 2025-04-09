package service

import (
	"observer-go/src/repositories"
	"observer-go/src/structs/DTO"
	"observer-go/src/structs/model"

	"gorm.io/gorm"
)

type CompanyServiceInterface interface {
	CreateCompany(dto DTO.CompanyDTO) error
	GetCompanyByID(id uint) (DTO.CompanyDTO, error)
	UpdateCompany(dto DTO.CompanyDTO) error
	DeleteCompany(id uint) error
}

type CompanyService struct {
	CompanyRepository repositories.CompanyRepoInterface
}

func NewCompanyService(companyRepository repositories.CompanyRepoInterface) *CompanyService {
	return &CompanyService{
		CompanyRepository: companyRepository,
	}
}

func (s *CompanyService) CreateCompany(dto DTO.CompanyDTO) error {
	company := model.Company{
		Name:    dto.Name,
		Address: dto.Address,
		Phone:   dto.Phone,
		Email:   dto.Email,
		OwnerID: dto.OwnerID,
	}
	return s.CompanyRepository.Create(company)
}

func (s *CompanyService) GetCompanyByID(id uint) (DTO.CompanyDTO, error) {
	company, err := s.CompanyRepository.GetById(id)
	if err != nil {
		return DTO.CompanyDTO{}, err
	}
	dto := DTO.CompanyDTO{
		ID:      company.ID,
		Name:    company.Name,
		Address: company.Address,
		Phone:   company.Phone,
		Email:   company.Email,
		OwnerID: company.OwnerID,
	}
	return dto, nil
}

func (s *CompanyService) UpdateCompany(dto DTO.CompanyDTO) error {
	company := model.Company{
		Model:   gorm.Model{ID: dto.ID},
		Name:    dto.Name,
		Address: dto.Address,
		Phone:   dto.Phone,
		Email:   dto.Email,
		OwnerID: dto.OwnerID,
	}
	return s.CompanyRepository.Update(company)
}

func (s *CompanyService) DeleteCompany(id uint) error {
	return s.CompanyRepository.Delete(id)
}
