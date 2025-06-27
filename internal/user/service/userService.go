package service

import (
	"goNext/internal/user/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository `inject:"type"`
}

func (us *UserService) GetUsers() ([]string, error) {
	if us.UserRepository == nil {
		panic("UserService.userRepository is nil! Dependency injection failed.")
	}
	users, error := us.UserRepository.FindAll()
	if error != nil {
		return nil, error
	}
	return users, nil
}
