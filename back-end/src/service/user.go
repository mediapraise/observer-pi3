package service

import (
	"observer-go/src/repositories"
	"observer-go/src/structs/DTO"
	"observer-go/src/structs/model"

	"gorm.io/gorm"
)

type UserServiceInterface interface {
	CreateUser(dto DTO.UserDTO) error
	GetUserByID(id uint) (DTO.UserDTO, error)
	// GetUserByEmail(email string) (DTO.UserDTO, error)
	GetAllUsers() ([]DTO.UserDTO, error)
	UpdateUser(dto DTO.UserDTO) error
	DeleteUser(id uint) error
}

type UserService struct {
	UserRepository repositories.UserRepoInterface
}

func NewUserService(userRepository repositories.UserRepoInterface) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) CreateUser(dto DTO.UserDTO) error {
	// Convert DTO to model and save to database
	user := model.User{
		Name:      dto.Name,
		Email:     dto.Email,
		Verified:  dto.Verified,
		IsAdmin:   dto.IsAdmin,
		CompanyID: dto.CompanyID,
	}
	return s.UserRepository.Create(user)

}

func (s *UserService) GetUserByID(id uint) (DTO.UserDTO, error) {
	// Retrieve user by ID from database
	user, err := s.UserRepository.GetById(id)
	if err != nil {
		return DTO.UserDTO{}, err
	}

	dto := DTO.UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Verified:  user.Verified,
		IsAdmin:   user.IsAdmin,
		CompanyID: user.CompanyID,
	}
	return dto, nil
}

// func (s *UserService) GetUserByEmail(email string) (DTO.UserDTO, error) {
// 	// Retrieve user by email from database
// 	user, err := s.UserRepository.GetByEmail(email)
// 	if err != nil {
// 		return DTO.UserDTO{}, err
// 	}
// 	dto := DTO.UserDTO{
// 		ID:        user.ID,
// 		Name:      user.Name,
// 		Email:     user.Email,
// 		Verified:  user.Verified,
// 		IsAdmin:   user.IsAdmin,
// 		CompanyID: user.CompanyID,
// 	}
// 	return dto, nil
// }

func (s *UserService) GetAllUsers() ([]DTO.UserDTO, error) {
	// Retrieve all users from database
	users, err := s.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var dtos []DTO.UserDTO
	for _, user := range users {
		dto := DTO.UserDTO{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Verified:  user.Verified,
			IsAdmin:   user.IsAdmin,
			CompanyID: user.CompanyID,
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}

func (s *UserService) UpdateUser(dto DTO.UserDTO) error {
	// Update user in database
	user := model.User{
		Model:     gorm.Model{ID: dto.ID},
		Name:      dto.Name,
		Email:     dto.Email,
		Verified:  dto.Verified,
		IsAdmin:   dto.IsAdmin,
		CompanyID: dto.CompanyID,
	}
	return s.UserRepository.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	// Delete user from database
	return s.UserRepository.Delete(id)
}
