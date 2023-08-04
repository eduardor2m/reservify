package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/primary"
)

type UserHandler struct {
	service primary.UserManager
}

func (instance UserHandler) CreateUser(context echo.Context) error {

	userCreated, err := user.NewBuilder().WithName("test").WithEmail("test@gmail.com").WithPassword("test").Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	err = instance.service.CreateUser(*userCreated)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, nil)
}

func (instance UserHandler) GetUserByName(context echo.Context) error {
	return context.JSON(http.StatusOK, nil)
}

func (instance UserHandler) UpdateUserByEmail(context echo.Context) error {
	return context.JSON(http.StatusOK, nil)
}

func (instance UserHandler) DeleteUserByEmail(context echo.Context) error {
	return context.JSON(http.StatusOK, nil)
}

func NewUserHandler(service primary.UserManager) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
