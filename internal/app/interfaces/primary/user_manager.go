package primary

import "reservify/internal/app/entity/user"

type UserManager interface {
	CreateUser(user user.User) error
	ListAllUsers() ([]user.User, error)
	GetUserByName(name string) ([]user.User, error)
	UpdateUserByEmail(email string, user user.User) error
	DeleteUserByEmail(email string) error
}
