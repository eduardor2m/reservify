package services

import (
	"github.com/google/uuid"
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/interfaces/repository"
)

var _ primary.UserManager = (*UserServices)(nil)

type UserServices struct {
	userRepository repository.UserLoader
}

func (instance *UserServices) CreateUser(u user.User) error {
	newUserUUID, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	formattedUser, err := user.NewBuilder().WithID(newUserUUID).WithName(u.Name()).WithEmail(u.Email()).WithPassword(u.Password()).WithDateOfBirth(u.DateOfBirth()).WithAdmin(u.Admin()).Build()

	if err != nil {
		return err
	}

	return instance.userRepository.CreateUser(*formattedUser)
}

func (instance *UserServices) ListAllUsers() ([]user.User, error) {
	return instance.userRepository.ListAllUsers()
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
