package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "reservify/internal/adapters/delivery/docs"
)

type Routes interface {
	Load(group *echo.Group)
}

type routes struct{}

func New() Routes {
	return &routes{}
}

func (instance *routes) Load(group *echo.Group) {
	group.GET("/docs/*", echoSwagger.WrapHandler)

	loadUserRoutes(group)
	loadRoomRoutes(group)
}
