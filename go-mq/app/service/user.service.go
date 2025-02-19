package service

import (
	"math"

	"github.com/bantawa04/go-mq/app/errors"
	"github.com/bantawa04/go-mq/app/model"
	"github.com/bantawa04/go-mq/app/repository"
	"github.com/bantawa04/go-mq/app/response"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers(page, perPage int, searchQuery string) ([]model.UserModel, *response.PaginationMeta, error)
	CreateUser(user *model.UserModel) error
	GetUserByEmail(email string) bool
	GetUserByPhone(phone string) bool
	GetUserById(id string) (*model.UserModel, error)
	UpdateUser(id string, user *model.UserModel) (*model.UserModel, error)
	DeleteUser(id string) error
	WithTrx(tx *gorm.DB) UserService
}

func (s *userService) WithTrx(tx *gorm.DB) UserService {
	newService := &userService{
		userRepo: s.userRepo.WithTrx(tx),
	}
	return newService
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) GetUserByEmail(email string) bool {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) ExistsByPhone(phone string) bool {
	return s.userRepo.GetUserByPhone(phone)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUsers(page, perPage int, searchQuery string) ([]model.UserModel, *response.PaginationMeta, error) {
	users, total, err := s.userRepo.GetUsers(page, perPage, searchQuery)
	if err != nil {
		return nil, nil, errors.NewInternalError(err)
	}

	if len(users) == 0 {
		return users, &response.PaginationMeta{
			Page:       page,
			PerPage:    perPage,
			TotalPages: 0,
			TotalItems: 0,
		}, nil
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	meta := &response.PaginationMeta{
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
		TotalItems: total,
	}

	return users, meta, nil
}

func (s *userService) GetUserByPhone(phone string) bool {
	return s.userRepo.GetUserByPhone(phone)
}

func (s *userService) CreateUser(user *model.UserModel) error {
	if user == nil {
		return errors.NewBadRequestError("User data cannot be empty")
	}

	if s.GetUserByEmail(user.Email) {
		return errors.NewConflictError("Email already exists")
	}

	err := s.userRepo.CreateUser(user)
	if err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}

func (s *userService) GetUserById(id string) (*model.UserModel, error) {
	if id == "" {
		return nil, errors.NewBadRequestError("User ID cannot be empty")
	}

	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if user == nil {
		return nil, errors.NewNotFoundError("User not found")
	}
	return user, nil
}

func (s *userService) UpdateUser(id string, updateData *model.UserModel) (*model.UserModel, error) {
	existingUser, err := s.userRepo.GetUserById(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if existingUser == nil {
		return nil, errors.NewNotFoundError("User not found")
	}

	if updateData.Email != "" && updateData.Email != existingUser.Email {
		if s.GetUserByEmail(updateData.Email) {
			return nil, errors.NewConflictError("Email already in use")
		}
	}

	if updateData.FullName != "" {
		existingUser.FullName = updateData.FullName
	}
	if *updateData.Phone != "" {
		existingUser.Phone = updateData.Phone
	}
	if *updateData.Gender != "" {
		existingUser.Gender = updateData.Gender
	}
	if updateData.Email != "" {
		existingUser.Email = updateData.Email
	}

	return s.userRepo.UpdateUser(existingUser)
}

func (s *userService) DeleteUser(id string) error {
	existingUser, err := s.userRepo.GetUserById(id)
	if err != nil {
		return errors.NewInternalError(err)
	}
	if existingUser == nil {
		return errors.NewNotFoundError("User not found")
	}

	return s.userRepo.DeleteUser(id)
}
