package repositories

import (
	"observer-go/src/db"
	"observer-go/src/structs/model"
)

type EventRepoInterface interface {
	GetAll() ([]model.Event, error)
	GetById(id uint) (model.Event, error)
	Create(event model.Event) error
	Update(event model.Event) error
	Delete(id uint) error
}

type EventRepo struct {
	db *db.Database
}

func NewEventRepo(db *db.Database) EventRepoInterface {
	return &EventRepo{
		db: db,
	}
}

func (repo *EventRepo) GetAll() ([]model.Event, error) {
	var events []model.Event
	err := repo.db.Gorm.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (repo *EventRepo) GetById(id uint) (model.Event, error) {
	var event model.Event
	err := repo.db.Gorm.First(&event, id).Error
	if err != nil {
		return model.Event{}, err
	}
	return event, nil
}

func (repo *EventRepo) Create(event model.Event) error {
	err := repo.db.Gorm.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *EventRepo) Update(event model.Event) error {
	err := repo.db.Gorm.Save(&event).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *EventRepo) Delete(id uint) error {
	var event model.Event
	err := repo.db.Gorm.First(&event, id).Error
	if err != nil {
		return err
	}
	err = repo.db.Gorm.Delete(&model.Event{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
