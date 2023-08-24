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

// @Summary Cria um novo usuário
// @Description Cria um novo usuário com base nos detalhes fornecidos
// @Tags Usuário
// @Produce json
// @Param userDTO body request.UserDTO true "Detalhes do usuário a ser criado"
// @Success 200 {object} response.InfoResponse "Usuário criado com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao criar usuário"
// @Router /users [post]
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

// @Summary Realiza o login de um usuário
// @Description Realiza o login de um usuário com base no email e senha fornecidos
// @Tags Usuário
// @Produce json
// @Param userDTO body request.UserDTO true "Detalhes do usuário para login"
// @Success 200 {object} response.InfoResponse "Usuário logado com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao realizar login"
// @Router /users/login [post]
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

// @Summary Lista todos os usuários
// @Description Retorna uma lista de todos os usuários registrados
// @Tags Usuário
// @Produce json
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {array} response.User
// @Failure 400 {object} response.ErrorResponse "Erro ao listar usuários"
// @Router /users [get]
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

// @Summary Obtém detalhes de um usuário por ID
// @Description Retorna os detalhes de um usuário com base no ID fornecido
// @Tags Usuário
// @Produce json
// @Param id path string true "ID do usuário"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {object} response.User
// @Failure 400 {object} response.ErrorResponse "Erro ao obter detalhes do usuário"
// @Router /users/{id} [get]
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

// @Summary Obtém usuários por nome
// @Description Retorna uma lista de usuários com base no nome fornecido
// @Tags Usuário
// @Produce json
// @Param name path string true "Nome do usuário"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {array} response.User
// @Failure 400 {object} response.ErrorResponse "Erro ao obter usuários por nome"
// @Router /users/name/{name} [get]
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

// @Summary Atualiza um usuário por email
// @Description Atualiza informações de um usuário com base no email fornecido
// @Tags Usuário
// @Produce json
// @Param email path string true "Email do usuário"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 423 {object} response.InfoResponse "Atualização de usuário não implementada ainda"
// @Router /users/{email} [put]
func (instance UserHandler) UpdateUserByEmail(context echo.Context) error {
	return context.JSON(http.StatusLocked, response.InfoResponse{Message: "Not implemented yet"})
}

// @Summary Atualiza o status de administrador de um usuário por ID
// @Description Atualiza o status de administrador de um usuário com base no ID fornecido
// @Tags Usuário
// @Produce json
// @Param id_user path string true "ID do usuário"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {object} response.InfoResponse "Campo admin do usuário foi atualizado"
// @Failure 400 {object} response.ErrorResponse "Erro ao atualizar status de administrador do usuário"
// @Router /users/admin/{id_user} [put]
func (instance UserHandler) UpdateAdminByUserID(context echo.Context) error {
	id := context.Param("id_user")
	idParse, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	token := context.Request().Header.Get("Authorization")

	err = instance.service.UpdateAdminByUserID(idParse, token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}
	
	return context.JSON(http.StatusOK, response.InfoResponse{Message: "campo admin do usuario foi atualizado"})
}

// @Summary Deleta um usuário por email
// @Description Deleta um usuário com base no email fornecido
// @Tags Usuário
// @Produce json
// @Param email path string true "Email do usuário"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {object} response.InfoResponse "Usuário deletado com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao deletar usuário"
// @Router /users/email/{email} [delete]
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
