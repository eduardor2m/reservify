package primary

import (
	"reservify/internal/app/entity/user"

	"github.com/google/uuid"
)

type UserManager interface {
	CreateUser(user user.User) error
	LoginUser(email string, password string) (*string, error)
	ListAllUsers(tokenJwt string) ([]user.User, error)
	GetUserByID(id uuid.UUID, tokenJwt string) (*user.User, error)
	GetUsersByName(name string, tokenJwt string) ([]user.User, error)
	UpdateUserByEmail(email string, tokenJwt string, user user.User) error
	DeleteUserByEmail(email string, tokenJwt string) error
}
