package routes

import (
	"github.com/labstack/echo/v4"
	"reservify/internal/adapters/delivery/http/dicontainer"
)

func loadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user")
	userHandlers := dicontainer.GetUserHandler()

	userGroup.POST("", userHandlers.CreateUser)
	userGroup.POST("/login", userHandlers.LoginUser)
	userGroup.POST("/rent", userHandlers.RentRoom)
	userGroup.GET("", userHandlers.ListAllUsers)
	userGroup.GET("/id/:id", userHandlers.GetUserByID)
	userGroup.GET("/:name", userHandlers.GetUserByName)
	userGroup.PUT("/:email", userHandlers.UpdateUserByEmail)
	userGroup.DELETE("/:email", userHandlers.DeleteUserByEmail)
}
