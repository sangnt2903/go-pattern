package services

import (
	"go-pattern/models"
	"go-pattern/repositories"
)

type userServiceInterface interface {
	FetchUsers(users *[]models.User) error
	StoreUser(user *models.User) error
}

type UserService struct {
	userRepository *repositories.UserRepository
}

func (us *UserService) StoreUser(user *models.User) error {
	return us.userRepository.Store(user)
}

func (us *UserService) FetchUsers(users *[]models.User) error {
	return us.userRepository.FetchAll(&users, "")
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo}
}

var _ userServiceInterface = &UserService{}
