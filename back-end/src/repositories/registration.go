package repositories

import (
	"observer-go/src/db"
	"observer-go/src/structs/model"
)

type RegistrationRepoInterface interface {
	GetAll() ([]model.Registration, error)
	GetById(id uint) (model.Registration, error)
	Create(registration model.Registration) error
	Update(registration model.Registration) error
	Delete(id uint) error
}

type RegistrationRepo struct {
	db *db.Database
}

func NewRegistrationRepo(db *db.Database) RegistrationRepoInterface {
	return &RegistrationRepo{
		db: db,
	}
}

func (repo *RegistrationRepo) GetAll() ([]model.Registration, error) {
	var registrations []model.Registration
	err := repo.db.Gorm.Find(&registrations).Error
	if err != nil {
		return nil, err
	}
	return registrations, nil
}

func (repo *RegistrationRepo) GetById(id uint) (model.Registration, error) {
	var registration model.Registration
	err := repo.db.Gorm.First(&registration, id).Error
	if err != nil {
		return model.Registration{}, err
	}
	return registration, nil
}

func (repo *RegistrationRepo) Create(registration model.Registration) error {
	err := repo.db.Gorm.Create(&registration).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *RegistrationRepo) Update(registration model.Registration) error {
	err := repo.db.Gorm.Save(&registration).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *RegistrationRepo) Delete(id uint) error {
	var registration model.Registration
	err := repo.db.Gorm.First(&registration, id).Error
	if err != nil {
		return err
	}
	err = repo.db.Gorm.Delete(&model.Registration{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
