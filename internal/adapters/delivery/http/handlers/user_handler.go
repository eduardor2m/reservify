package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservify/internal/adapters/delivery/http/handlers/dto/request"
	"reservify/internal/adapters/delivery/http/handlers/dto/response"
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/primary"
)

type UserHandler struct {
	service primary.UserManager
}

func (instance UserHandler) CreateUser(context echo.Context) error {
	var userDTO request.UserDTO

	err := context.Bind(&userDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	userReceived, err := user.NewBuilder().WithName(userDTO.Name).WithEmail(userDTO.Email).WithPassword(userDTO.Password).WithDateOfBirth(userDTO.DateOfBirth).WithAdmin(userDTO.Admin).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.CreateUser(*userReceived)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User created successfully"})
}

func (instance UserHandler) LoginUser(context echo.Context) error {
	var userDTO request.UserDTO

	err := context.Bind(&userDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err, token := instance.service.LoginUser(userDTO.Email, userDTO.Password)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	context.Response().Header().Set("Authorization", *token)
	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User logged successfully"})
}

func (instance UserHandler) ListAllUsers(context echo.Context) error {
	users, err := instance.service.ListAllUsers()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var usersResponse []response.User

	for _, userDB := range users {
		usersResponse = append(usersResponse, *response.NewUser(userDB))
	}

	return context.JSON(http.StatusOK, usersResponse)
}

func (instance UserHandler) GetUserByName(context echo.Context) error {
	var name string

	name = context.Param("name")

	users, err := instance.service.GetUserByName(name)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var usersResponse []response.User

	for _, userDB := range users {
		usersResponse = append(usersResponse, *response.NewUser(userDB))
	}

	return context.JSON(http.StatusOK, usersResponse)

}

func (instance UserHandler) UpdateUserByEmail(context echo.Context) error {
	return context.JSON(http.StatusLocked, response.InfoResponse{Message: "Not implemented yet"})
}

func (instance UserHandler) DeleteUserByEmail(context echo.Context) error {
	var email string

	email = context.Param("email")

	err := instance.service.DeleteUserByEmail(email)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User deleted successfully"})
}

func NewUserHandler(service primary.UserManager) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
