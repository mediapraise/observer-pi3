package service

import (
	"observer-go/src/repositories"
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
	ClientRepository repositories.ClientRepoInterface
}

func NewClientService(clientRepository repositories.ClientRepoInterface)  *ClientService {
	return &ClientService{
		ClientRepository: clientRepository,
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
	return s.ClientRepository.Create(client)
}

func (s *ClientService) GetClientByID(id uint) (DTO.ClientDTO, error) {
	client, err := s.ClientRepository.GetById(id)
	if err != nil {
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
	return s.ClientRepository.Update(client)
}

func (s *ClientService) DeleteClient(id uint) error {
	return s.ClientRepository.Delete(id)
}
