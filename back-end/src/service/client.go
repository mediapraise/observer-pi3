package service

import (
	"observer-go/src/db"
	"observer-go/src/structs/DTO"
	"observer-go/src/structs/model"

	"gorm.io/gorm"
)

type ClientServiceInterface interface {
	CreateClient(dto DTO.ClientDTO) error
	GetClientByID(id uint) (DTO.ClientDTO, error)
	UpdateClient(dto DTO.ClientDTO) error
	DeleteClient(id uint) error
}

type ClientService struct {
	Database *db.Database
}

func NewClientService(database *db.Database) *ClientService {
	return &ClientService{
		Database: database,
	}
}

func (s *ClientService) CreateClient(dto DTO.ClientDTO) error {
	client := model.Client{
		Name:      dto.Name,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Document:  dto.Document,
		CompanyID: dto.CompanyID,
	}
	return s.Database.Gorm.Create(&client).Error
}

func (s *ClientService) GetClientByID(id uint) (DTO.ClientDTO, error) {
	var client model.Client
	if err := s.Database.Gorm.First(&client, id).Error; err != nil {
		return DTO.ClientDTO{}, err
	}
	dto := DTO.ClientDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		Document:  client.Document,
		CompanyID: client.CompanyID,
	}
	return dto, nil
}

func (s *ClientService) UpdateClient(dto DTO.ClientDTO) error {
	client := model.Client{
		Model:     gorm.Model{ID: dto.ID},
		Name:      dto.Name,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Document:  dto.Document,
		CompanyID: dto.CompanyID,
	}
	return s.Database.Gorm.Save(&client).Error
}

func (s *ClientService) DeleteClient(id uint) error {
	return s.Database.Gorm.Delete(&model.Client{}, id).Error
}
