package services

import (
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/interfaces/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password()), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	encryptedPasswordString := string(encryptedPassword)

	date := time.Now()

	formattedUser, err := user.NewBuilder().WithID(newUserUUID).WithName(u.Name()).WithCPF(u.CPF()).WithDateOfBirth(u.DateOfBirth()).WithPhone(u.Phone()).WithEmail(u.Email()).WithPassword(encryptedPasswordString).WithAdmin(u.Admin()).WithCreatedAt(date).WithUpdatedAt(date).Build()

	if err != nil {
		return err
	}

	return instance.userRepository.CreateUser(*formattedUser)
}

func (instance *UserServices) LoginUser(email string, password string) (*string, error) {
	return instance.userRepository.LoginUser(email, password)
}

func (instance *UserServices) ListAllUsers(tokenJwt string) ([]user.User, error) {
	return instance.userRepository.ListAllUsers(tokenJwt)
}

func (instance *UserServices) GetUserByID(id uuid.UUID) (*user.User, error) {
	return instance.userRepository.GetUserByID(id)
}

func (instance *UserServices) GetUsersByName(name string) ([]user.User, error) {
	return instance.userRepository.GetUsersByName(name)
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
