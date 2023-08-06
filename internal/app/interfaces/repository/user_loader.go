package repository

import (
	"reservify/internal/app/entity/guest"
	"reservify/internal/app/entity/user"
)

type UserLoader interface {
	CreateUser(user user.User) error
	LoginUser(email string, password string) (error, *string)
	RentRoom(cpf string, roomCod string, guests *[]guest.Guest) error
	ListAllUsers() ([]user.User, error)
	GetUserByName(name string) ([]user.User, error)
	UpdateUserByEmail(email string, user user.User) error
	DeleteUserByEmail(email string) error
}
