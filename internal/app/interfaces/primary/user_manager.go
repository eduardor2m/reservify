package primary

import (
	"reservify/internal/app/entity/reservation"
	"reservify/internal/app/entity/user"

	"github.com/google/uuid"
)

type UserManager interface {
	CreateUser(user user.User) error
	LoginUser(email string, password string) (*string, error)
	CreateReservation(reservation reservation.Reservation) error
	GetReservationByID(id uuid.UUID) (*reservation.Reservation, error)
	GetReservationByIDRoom(idRoom uuid.UUID) ([]reservation.Reservation, error)
	GetReservationByIDUser(idUser uuid.UUID) ([]reservation.Reservation, error)
	DeleteReservationByID(id uuid.UUID) error
	ListAllReservations() ([]reservation.Reservation, error)
	ListAllUsers() ([]user.User, error)
	GetUserByID(id uuid.UUID) (*user.User, error)
	GetUserByName(name string) ([]user.User, error)
	UpdateUserByEmail(email string, user user.User) error
	DeleteUserByEmail(email string) error
}
