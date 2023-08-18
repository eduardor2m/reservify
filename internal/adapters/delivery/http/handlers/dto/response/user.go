package response

import (
	"github.com/google/uuid"
	"reservify/internal/app/entity/user"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CPF         string    `json:"cpf"`
	Phone       string    `json:"phone"`
	DateOfBirth string    `json:"date_of_birth"`
	Admin       bool      `json:"admin"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

func NewUser(user user.User) *User {
	return &User{
		ID:          user.ID(),
		Name:        user.Name(),
		Email:       user.Email(),
		CPF:         user.CPF(),
		Phone:       user.Phone(),
		DateOfBirth: user.DateOfBirth().String(),
		Admin:       user.Admin(),
		CreatedAt:   user.CreatedAt().String(),
		UpdatedAt:   user.UpdatedAt().String(),
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type InfoResponse struct {
	Message string `json:"message"`
}
