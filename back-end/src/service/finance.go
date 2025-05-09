package service

import (
	"observer-go/src/repositories"
	"observer-go/src/structs/DTO"
	"observer-go/src/structs/model"
	"strconv"

	"gorm.io/gorm"
)

type FinanceServiceInterface interface {
	Create(dto DTO.HistoryPaymentDTO) error
	GetByID(id uint) (DTO.HistoryPaymentDTO, error)
	GetAllByCompanyID(companyID uint) ([]DTO.HistoryPaymentDTO, error)
	GetAll() ([]DTO.HistoryPaymentDTO, error)
	Update(dto DTO.HistoryPaymentDTO) error
}
type HistoryPaymentServiceInterface interface {
	FinanceServiceInterface
}

func parseUint(value string) uint {
	parsedValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		panic("Invalid CompanyID format: " + err.Error())
	}
	return uint(parsedValue)
}

type HistoryPaymentService struct {
	HistoryPaymentRepository repositories.HistoryPaymentRepoInterface
}

func NewHistoryPaymentService(historyPaymentRepo repositories.HistoryPaymentRepoInterface) *HistoryPaymentService {
	return &HistoryPaymentService{HistoryPaymentRepository: historyPaymentRepo}
}

func (s *HistoryPaymentService) Create(dto DTO.HistoryPaymentDTO) error {
	historyPayment := model.HistoryPayment{
		CompanyID:     parseUint(dto.CompanyID),
		PaymentAmount: dto.PaymentAmount,
		PaymentDate:   dto.PaymentDate,
	}
	return s.HistoryPaymentRepository.Create(historyPayment)
}

func (s *HistoryPaymentService) GetByID(id uint) (DTO.HistoryPaymentDTO, error) {
	historyPayment, err := s.HistoryPaymentRepository.GetById(id)
	if err != nil {
		return DTO.HistoryPaymentDTO{}, err
	}
	dto := DTO.HistoryPaymentDTO{
		ID:            historyPayment.ID,
		CompanyID:     strconv.FormatUint(uint64(historyPayment.CompanyID), 10),
		PaymentAmount: historyPayment.PaymentAmount,
		PaymentDate:   historyPayment.PaymentDate,
	}
	return dto, nil
}

func (s *HistoryPaymentService) GetAllByCompanyID(companyID uint) ([]DTO.HistoryPaymentDTO, error) {
	historyPayments, err := s.HistoryPaymentRepository.GetAllByCompanyID(companyID)
	if err != nil {
		return nil, err
	}
	var dtos []DTO.HistoryPaymentDTO
	for _, historyPayment := range historyPayments {
		dtos = append(dtos, DTO.HistoryPaymentDTO{
			ID:            historyPayment.ID,
			CompanyID:     strconv.FormatUint(uint64(historyPayment.CompanyID), 10),
			PaymentAmount: historyPayment.PaymentAmount,
			PaymentDate:   historyPayment.PaymentDate,
		})
	}
	return dtos, nil
}

func (s *HistoryPaymentService) GetAll() ([]DTO.HistoryPaymentDTO, error) {
	historyPayments, err := s.HistoryPaymentRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var dtos []DTO.HistoryPaymentDTO
	for _, historyPayment := range historyPayments {
		dtos = append(dtos, DTO.HistoryPaymentDTO{
			ID:            historyPayment.ID,
			CompanyID:     strconv.FormatUint(uint64(historyPayment.CompanyID), 10),
			PaymentAmount: historyPayment.PaymentAmount,
			PaymentDate:   historyPayment.PaymentDate,
		})
	}
	return dtos, nil
}

func (s *HistoryPaymentService) Update(dto DTO.HistoryPaymentDTO) error {
	historyPayment := model.HistoryPayment{
		Model:         gorm.Model{ID: dto.ID},
		CompanyID:     parseUint(dto.CompanyID),
		PaymentAmount: dto.PaymentAmount,
		PaymentDate:   dto.PaymentDate,
	}
	return s.HistoryPaymentRepository.Update(historyPayment)
}

func (s *HistoryPaymentService) Delete(id uint) error {
	return s.HistoryPaymentRepository.Delete(id)
}