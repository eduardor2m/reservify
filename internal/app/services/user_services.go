package services

import (
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/interfaces/repository"
)

var _ primary.UserManager = (*UserServices)(nil)

type UserServices struct {
	userRepository repository.UserLoader
}

func (instance *UserServices) CreateUser(user user.User) error {
	return instance.userRepository.CreateUser(user)
}

func (instance *UserServices) GetUserByName(name string) ([]user.User, error) {
	return instance.userRepository.GetUserByName(name)
}

func (instance *UserServices) UpdateUserByEmail(email string, user user.User) error {
	return instance.userRepository.UpdateUserByEmail(email, user)
}

func (instance *UserServices) DeleteUserByEmail(email string) error {
	return instance.userRepository.DeleteUserByEmail(email)
}

func NewUserServices(userRepository repository.UserLoader) *UserServices {
	return &UserServices{
		userRepository: userRepository,
	}
}
