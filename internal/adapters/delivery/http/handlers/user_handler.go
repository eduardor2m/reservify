package handlers

import (
	"net/http"
	"reservify/internal/adapters/delivery/http/handlers/dto/request"
	"reservify/internal/adapters/delivery/http/handlers/dto/response"
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/primary"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service primary.UserManager
}

// CreateUser godoc
// @ID CreateUser
// @Summary Cria um novo usuário.
// @Description Cria um novo usuário.
// @Security bearerAuth
// @Tags User
// @Produce json
// @Success 200 {object} response.InfoResponse "User created successfully"
// @Failure 401 {object} response.ErrorMessage "Acesso não autorizado."
// @Router /user [post]

func (instance UserHandler) CreateUser(context echo.Context) error {
	var userDTO request.UserDTO

	err := context.Bind(&userDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	userReceived, err := user.NewBuilder().WithName(userDTO.Name).WithEmail(userDTO.Email).WithCPF(userDTO.CPF).WithPhone(userDTO.Phone).WithPassword(userDTO.Password).WithDateOfBirth(userDTO.DateOfBirth).WithAdmin(userDTO.Admin).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.CreateUser(*userReceived)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User created successfully"})
}

// LoginUser
// @ID LoginUser
// @Summary Realiza o login do usuário.
// @Description Realiza o login do usuário.
// @Security bearerAuth
// @Tags User
// @Produce json
// @Success 200 {object} response.InfoResponse "User logged successfully"
// @Failure 401 {object} response.ErrorMessage
// @Router /user/login [post]

func (instance UserHandler) LoginUser(context echo.Context) error {
	var userDTO request.UserDTO

	err := context.Bind(&userDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	token, err := instance.service.LoginUser(userDTO.Email, userDTO.Password)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	context.Response().Header().Set("Authorization", *token)
	return context.JSON(http.StatusOK, response.InfoResponse{Message: "User logged successfully"})
}

// ListAllUsers
// @ID ListAllUsers
// @Summary Lista todos os usuários.
// @Description Lista todos os usuários.
// @Security bearerAuth
// @Tags User
// @Produce json
// @Success 200 {object} []User
// @Failure 401 {object} response.ErrorMessage
// @Router /user [get]

func (instance UserHandler) ListAllUsers(context echo.Context) error {
	token := context.Request().Header.Get("Authorization")

	users, err := instance.service.ListAllUsers(token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var usersResponse []response.User

	for _, userDB := range users {
		usersResponse = append(usersResponse, *response.NewUser(userDB))
	}

	if len(usersResponse) == 0 {
		return context.JSON(http.StatusOK, []response.User{})
	}

	return context.JSON(http.StatusOK, usersResponse)
}

// GetUserByID

func (instance UserHandler) GetUserByID(context echo.Context) error {
	id := context.Param("id")

	userID, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	token := context.Request().Header.Get("Authorization")

	userReceived, err := instance.service.GetUserByID(userID, token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.NewUser(*userReceived))
}

// GetUserByName
// @ID GetUserByName
// @Summary Busca um usuário pelo nome.
// @Description Busca um usuário pelo nome.
// @Security bearerAuth
// @Tags User
// @Produce json
// @Param name path string true "Nome do usuário"
// @Success 200 {object} []User
// @Failure 401 {object} response.ErrorMessage
// @Router /user/{name} [get]

func (instance UserHandler) GetUsersByName(context echo.Context) error {
	name := context.Param("name")
	token := context.Request().Header.Get("Authorization")

	users, err := instance.service.GetUsersByName(name, token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var usersResponse []response.User

	for _, userDB := range users {
		usersResponse = append(usersResponse, *response.NewUser(userDB))
	}

	if len(usersResponse) == 0 {
		return context.JSON(http.StatusOK, []response.User{})
	}

	return context.JSON(http.StatusOK, usersResponse)

}

// UpdateUserByEmail
// @ID UpdateUserByEmail
// @Summary Atualiza um usuário pelo email.
// @Description Atualiza um usuário pelo email.
// @Security bearerAuth
// @Tags User
// @Produce json
// @Param email path string true "Email do usuário"
// @Success 200 {object} []User
// @Failure 401 {object} response.ErrorMessage
// @Router /user/{email} [put]

func (instance UserHandler) UpdateUserByEmail(context echo.Context) error {
	return context.JSON(http.StatusLocked, response.InfoResponse{Message: "Not implemented yet"})
}

// DeleteUserByEmail
// @ID DeleteUserByEmail
// @Summary Deleta um usuário pelo email.
// @Description Deleta um usuário pelo email.
// @Security bearerAuth
// @Tags User
// @Produce json
// @Param email path string true "Email do usuário"
// @Success 200 {object} []User
// @Failure 401 {object} response.ErrorMessage
// @Router /user/{email} [delete]

func (instance UserHandler) DeleteUserByEmail(context echo.Context) error {
	email := context.Param("email")
	token := context.Request().Header.Get("Authorization")

	err := instance.service.DeleteUserByEmail(email, token)

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
