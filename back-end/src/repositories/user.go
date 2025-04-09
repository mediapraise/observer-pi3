package repositories

import (
	"observer-go/src/db"
	"observer-go/src/structs/model"
)

type UserRepoInterface interface {
	GetAll() ([]model.User, error)
	GetById(id uint) (model.User, error)
	Create(user model.User) error
	Update(user model.User) error
	Delete(id uint) error
}

type UserRepo struct {
	db *db.Database
}

func NewUserRepo(db *db.Database) UserRepoInterface {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) GetAll() ([]model.User, error) {
	var users []model.User
	err := repo.db.Gorm.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) GetById(id uint) (model.User, error) {
	var user model.User
	err := repo.db.Gorm.First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (repo *UserRepo) Create(user model.User) error {
	err := repo.db.Gorm.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) Update(user model.User) error {
	err := repo.db.Gorm.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) Delete(id uint) error {
	var user model.User
	err := repo.db.Gorm.First(&user, id).Error
	if err != nil {
		return err
	}
	err = repo.db.Gorm.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
