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

func (instance *UserServices) GetUserByID(id uuid.UUID, tokenJwt string) (*user.User, error) {
	return instance.userRepository.GetUserByID(id, tokenJwt)
}

func (instance *UserServices) GetUsersByName(name string, tokenJwt string) ([]user.User, error) {
	return instance.userRepository.GetUsersByName(name, tokenJwt)
}

func (instance *UserServices) UpdateUserByEmail(email string, tokenJwt string, user user.User) error {
	return instance.userRepository.UpdateUserByEmail(email, tokenJwt, user)
}

func (instance *UserServices) UpdateAdminByUserID(userID uuid.UUID, tokenJwt string) error {
	return instance.userRepository.UpdateAdminByUserID(userID, tokenJwt)
}

func (instance *UserServices) DeleteUserByEmail(email string, tokenJwt string) error {
	return instance.userRepository.DeleteUserByEmail(email, tokenJwt)
}

func NewUserServices(userRepository repository.UserLoader) *UserServices {
	return &UserServices{
		userRepository: userRepository,
	}
}
