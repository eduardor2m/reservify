package routes

import (
	"github.com/labstack/echo/v4"
	"reservify/internal/adapters/delivery/http/dicontainer"
)

func loadRoomRoutes(group *echo.Group) {
	roomGroup := group.Group("/room")
	roomHandlers := dicontainer.GetRoomHandler()

	roomGroup.POST("", roomHandlers.CreateRoom)
	roomGroup.GET("", roomHandlers.ListAllRooms)
	roomGroup.GET("/:id", roomHandlers.GetRoomByID)
}
