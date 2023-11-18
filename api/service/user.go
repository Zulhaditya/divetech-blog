package service

import (
	"models"
	"repository"
)

type UserService struct {
	repo repository.UserRepository
}

// get injected user repo
func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

// save users entity
func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

// get validated user
func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)
}
