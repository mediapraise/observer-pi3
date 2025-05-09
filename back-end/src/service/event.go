package service

import (
	"observer-go/src/repositories"
	"observer-go/src/structs/DTO"
	"observer-go/src/structs/model"

	"gorm.io/gorm"
)

type EventServiceInterface interface {
	CreateEvent(dto DTO.EventDTO) error
	GetEventByID(id uint) (DTO.EventDTO, error)
	GetAllEvents() ([]DTO.EventDTO, error)
	UpdateEvent(dto DTO.EventDTO) error
	DeleteEvent(id uint) error
}

type EventService struct {
	EventRepository repositories.EventRepoInterface
}

func NewEventService(eventRepo repositories.EventRepoInterface) *EventService {
	return &EventService{EventRepository: eventRepo}
}

func (s *EventService) CreateEvent(dto DTO.EventDTO) error {
	event := model.Event{
		Date:      dto.Date,
		Board:     dto.Board,
		CompanyID: dto.CompanyID,
		Event:     dto.Event,
	}
	return s.EventRepository.Create(event)
}

func (s *EventService) GetEventByID(id uint) (DTO.EventDTO, error) {
	event, err := s.EventRepository.GetById(id)
	if err != nil {
		return DTO.EventDTO{}, err
	}
	dto := DTO.EventDTO{
		ID:        event.ID,
		Date:      event.Date,
		Board:     event.Board,
		CompanyID: event.CompanyID,
		Event:     event.Event,
	}
	return dto, nil
}

func (s *EventService) GetAllEvents() ([]DTO.EventDTO, error) {
	events, err := s.EventRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var dtos []DTO.EventDTO
	for _, event := range events {
		dtos = append(dtos, DTO.EventDTO{
			ID:        event.ID,
			Date:      event.Date,
			Board:     event.Board,
			CompanyID: event.CompanyID,
			Event:     event.Event,
		})
	}
	return dtos, nil
}

func (s *EventService) UpdateEvent(dto DTO.EventDTO) error {
	event := model.Event{
		Model:     gorm.Model{ID: dto.ID},
		Date:      dto.Date,
		Board:     dto.Board,
		CompanyID: dto.CompanyID,
		Event:     dto.Event,
	}
	return s.EventRepository.Update(event)
}

func (s *EventService) DeleteEvent(id uint) error {
	return s.EventRepository.Delete(id)
}
