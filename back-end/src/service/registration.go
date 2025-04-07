package service

import (
	"observer-go/src/db"
	"observer-go/src/structs/DTO"
	"observer-go/src/structs/model"

	"gorm.io/gorm"
)

type RegistrationServiceInterface interface {
	CreateRegistration(dto DTO.RegistrationDTO) error
	GetRegistrationByID(id uint) (DTO.RegistrationDTO, error)
	UpdateRegistration(dto DTO.RegistrationDTO) error
	DeleteRegistration(id uint) error
}

type RegistrationService struct {
	Database *db.Database
}

func NewRegistrationService(database *db.Database) *RegistrationService {
	return &RegistrationService{
		Database: database,
	}
}

func (s *RegistrationService) CreateRegistration(dto DTO.RegistrationDTO) error {
	registration := model.Registration{
		Board:              dto.Board,
		Vehicle:            dto.Vehicle,
		ExpirationDate:     dto.ExpirationDate,
		RegistrationStatus: dto.RegistrationStatus,
		Owner:              dto.Owner,
		CompanyID:          dto.CompanyID,
	}
	return s.Database.Gorm.Create(&registration).Error
}

func (s *RegistrationService) GetRegistrationByID(id uint) (DTO.RegistrationDTO, error) {
	var registration model.Registration
	if err := s.Database.Gorm.First(&registration, id).Error; err != nil {
		return DTO.RegistrationDTO{}, err
	}
	dto := DTO.RegistrationDTO{
		ID:                 registration.ID,
		Board:              registration.Board,
		Vehicle:            registration.Vehicle,
		ExpirationDate:     registration.ExpirationDate,
		RegistrationStatus: registration.RegistrationStatus,
		Owner:              registration.Owner,
		CompanyID:          registration.CompanyID,
	}
	return dto, nil
}

func (s *RegistrationService) UpdateRegistration(dto DTO.RegistrationDTO) error {
	registration := model.Registration{
		Model:              gorm.Model{ID: dto.ID},
		Board:              dto.Board,
		Vehicle:            dto.Vehicle,
		ExpirationDate:     dto.ExpirationDate,
		RegistrationStatus: dto.RegistrationStatus,
		Owner:              dto.Owner,
		CompanyID:          dto.CompanyID,
	}
	return s.Database.Gorm.Save(&registration).Error
}

func (s *RegistrationService) DeleteRegistration(id uint) error {
	return s.Database.Gorm.Delete(&model.Registration{}, id).Error
}
