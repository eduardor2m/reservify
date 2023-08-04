package routes

import (
	"github.com/labstack/echo/v4"
	"reservify/internal/adapters/delivery/http/dicontainer"
)

func loadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user")
	userHandlers := dicontainer.GetUserHandler()

	userGroup.POST("", userHandlers.CreateUser)
}