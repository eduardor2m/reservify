package routes

import (
	"reservify/internal/adapters/delivery/http/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user")
	userHandlers := dicontainer.GetUserHandler()

	userGroup.POST("", userHandlers.CreateUser)
	userGroup.POST("/login", userHandlers.LoginUser)
	userGroup.POST("/reservation", userHandlers.CreateReservation)
	userGroup.GET("/reservation", userHandlers.ListAllReservations)
	userGroup.GET("", userHandlers.ListAllUsers)
	userGroup.GET("/id/:id", userHandlers.GetUserByID)
	userGroup.GET("/:name", userHandlers.GetUserByName)
	userGroup.PUT("/:email", userHandlers.UpdateUserByEmail)
	userGroup.DELETE("/:email", userHandlers.DeleteUserByEmail)
}
