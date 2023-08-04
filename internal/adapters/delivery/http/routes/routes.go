package routes

import "github.com/labstack/echo/v4"

type Routes interface {
	Load(group *echo.Group)
}

type routes struct{}

func New() Routes {
	return &routes{}
}

func (instance *routes) Load(group *echo.Group) {
	loadUserRoutes(group)
}
