package primary

import (
	"github.com/google/uuid"
	"reservify/internal/app/entity/user"
)

type UserManager interface {
	CreateUser(user user.User) error
	LoginUser(email string, password string) (error, *string)
	RentRoom(idUser string, idRoom string, checkIn string, checkOut string) error
	ListAllUsers() ([]user.User, error)
	GetUserByID(id uuid.UUID) (*user.User, error)
	GetUserByName(name string) ([]user.User, error)
	UpdateUserByEmail(email string, user user.User) error
	DeleteUserByEmail(email string) error
}
