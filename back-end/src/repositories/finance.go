package repositories

import (
	"observer-go/src/db"
	"observer-go/src/structs/model"
)

type FinanceRepoInterface interface {
	GetAll() ([]model.Event, error)
	GetAllByCompanyID(companyID uint) ([]model.Event, error)
	GetById(id uint) (model.Event, error)
	Create(event model.Event) error
	Update(event model.Event) error
	Delete(id uint) error
}

type FinanceRepo struct {
	db *db.Database
}

func NewFinanceRepo(db *db.Database) FinanceRepoInterface {
	return &FinanceRepo{
		db: db,
	}
}

func (repo *FinanceRepo) GetAll() ([]model.Event, error) {
	var events []model.Event
	err := repo.db.Gorm.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (repo *FinanceRepo) GetAllByCompanyID(companyID uint) ([]model.Event, error) {
	var events []model.Event
	err := repo.db.Gorm.Where("company_id = ?", companyID).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (repo *FinanceRepo) GetById(id uint) (model.Event, error) {
	var event model.Event
	err := repo.db.Gorm.First(&event, id).Error
	if err != nil {
		return model.Event{}, err
	}
	return event, nil
}

func (repo *FinanceRepo) Create(event model.Event) error {
	err := repo.db.Gorm.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *FinanceRepo) Update(event model.Event) error {
	err := repo.db.Gorm.Save(&event).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *FinanceRepo) Delete(id uint) error {
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
