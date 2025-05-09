package repositories

import (
	"observer-go/src/db"
	"observer-go/src/structs/model"
)

type HistoryPaymentRepoInterface interface {
	GetAll() ([]model.HistoryPayment, error)
	GetAllByCompanyID(companyID uint) ([]model.HistoryPayment, error)
	GetById(id uint) (model.HistoryPayment, error)
	Create(history model.HistoryPayment) error
	Update(history model.HistoryPayment) error
	Delete(id uint) error
}

type HistoryPaymentRepo struct {
	db *db.Database
}

func NewFinanceRepo(db *db.Database) HistoryPaymentRepoInterface {
	return &HistoryPaymentRepo{
		db: db,
	}
}

func (repo *HistoryPaymentRepo) GetAll() ([]model.HistoryPayment, error) {
	var historys []model.HistoryPayment
	err := repo.db.Gorm.Find(&historys).Error
	if err != nil {
		return nil, err
	}
	return historys, nil
}

func (repo *HistoryPaymentRepo) GetAllByCompanyID(companyID uint) ([]model.HistoryPayment, error) {
	var historys []model.HistoryPayment
	err := repo.db.Gorm.Where("company_id = ?", companyID).Find(&historys).Error
	if err != nil {
		return nil, err
	}
	return historys, nil
}

func (repo *HistoryPaymentRepo) GetById(id uint) (model.HistoryPayment, error) {
	var history model.HistoryPayment
	err := repo.db.Gorm.First(&history, id).Error
	if err != nil {
		return model.HistoryPayment{}, err
	}
	return history, nil
}

func (repo *HistoryPaymentRepo) Create(history model.HistoryPayment) error {
	err := repo.db.Gorm.Create(&history).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *HistoryPaymentRepo) Update(history model.HistoryPayment) error {
	err := repo.db.Gorm.Save(&history).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *HistoryPaymentRepo) Delete(id uint) error {
	var history model.Event
	err := repo.db.Gorm.First(&history, id).Error
	if err != nil {
		return err
	}
	err = repo.db.Gorm.Delete(&model.HistoryPayment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
