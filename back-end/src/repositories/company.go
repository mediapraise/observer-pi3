package repositories

import (
	"observer-go/src/db"
	"observer-go/src/structs/model"
)

type CompanyRepoInterface interface {
	GetAll() ([]model.Company, error)
	GetById(id uint) (model.Company, error)
	Create(company model.Company) error
	Update(company model.Company) error
	Delete(id uint) error
}

type CompanyRepo struct {
	db *db.Database
}

func NewCompanyRepo(db *db.Database) CompanyRepoInterface {
	return &CompanyRepo{
		db: db,
	}
}

func (repo *CompanyRepo) GetAll() ([]model.Company, error) {
	var companies []model.Company
	err := repo.db.Gorm.Find(&companies).Error
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (repo *CompanyRepo) GetById(id uint) (model.Company, error) {
	var company model.Company
	err := repo.db.Gorm.First(&company, id).Error
	if err != nil {
		return model.Company{}, err
	}
	return company, nil
}

func (repo *CompanyRepo) Create(company model.Company) error {
	err := repo.db.Gorm.Create(&company).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CompanyRepo) Update(company model.Company) error {
	err := repo.db.Gorm.Save(&company).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CompanyRepo) Delete(id uint) error {
	var company model.Company
	err := repo.db.Gorm.First(&company, id).Error
	if err != nil {
		return err
	}
	err = repo.db.Gorm.Delete(&model.Company{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
