package services

import (
	"reservify/internal/app/entity/reservation"
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

	formattedUser, err := user.NewBuilder().WithID(newUserUUID).WithName(u.Name()).WithCPF(u.CPF()).WithDateOfBirth(u.DateOfBirth()).WithPhone(u.Phone()).WithEmail(u.Email()).WithPassword(encryptedPasswordString).WithAdmin(u.Admin()).Build()

	if err != nil {
		return err
	}

	return instance.userRepository.CreateUser(*formattedUser)
}

func (instance *UserServices) LoginUser(email string, password string) (error, *string) {
	return instance.userRepository.LoginUser(email, password)
}

func (instance *UserServices) CreateReservation(
	r reservation.Reservation,
) error {
	reservationUUID, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	reservationTime := time.Now()

	reservationFormatted, err := reservation.NewBuilder().WithID(reservationUUID).WithIdRoom(r.IDRoom()).WithIdUser(r.IDUser()).WithCheckIn(r.CheckIn()).WithCheckOut(r.CheckOut()).WithCreatedAt(reservationTime).WithUpdatedAt(reservationTime).Build()

	if err != nil {
		return err
	}
	
	return instance.userRepository.CreateReservation(*reservationFormatted)
}

func (instance *UserServices) ListAllReservations() ([]reservation.Reservation, error) {
	return instance.userRepository.ListAllReservations()
}

func (instance *UserServices) ListAllUsers() ([]user.User, error) {
	return instance.userRepository.ListAllUsers()
}

func (instance *UserServices) GetUserByID(id uuid.UUID) (*user.User, error) {
	return instance.userRepository.GetUserByID(id)
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
