package repositories

import (
	"observer-go/src/db"
	"observer-go/src/structs/model"
)

type ClientRepoInterface interface {
	GetAll() ([]model.Client, error)
	GetById(id uint) (model.Client, error)
	Create(client model.Client) error
	Update(client model.Client) error
	Delete(id uint) error
}
type ClientRepo struct {
	db *db.Database
}

func NewClientRepo(db *db.Database) ClientRepoInterface {
	return &ClientRepo{
		db: db,
	}
}
func (repo *ClientRepo) GetAll() ([]model.Client, error) {
	var clients []model.Client
	err := repo.db.Gorm.Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (repo *ClientRepo) GetById(id uint) (model.Client, error) {
	var client model.Client
	err := repo.db.Gorm.First(&client, id).Error
	if err != nil {
		return model.Client{}, err
	}		
	return client, nil
}

func (repo *ClientRepo) Create(user model.Client) error {
	err := repo.db.Gorm.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ClientRepo) Update(user model.Client) error {
	err := repo.db.Gorm.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ClientRepo) Delete(id uint) error {
	var client model.Client
	err := repo.db.Gorm.First(&client, id).Error
	if err != nil {
		return err
	}
	err = repo.db.Gorm.Delete(&model.Client{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

