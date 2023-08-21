package routes

import (
	"reservify/internal/adapters/delivery/http/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadRoomRoutes(group *echo.Group) {
	roomGroup := group.Group("/room")
	roomHandlers := dicontainer.GetRoomHandler()

	roomGroup.POST("", roomHandlers.CreateRoom)
	roomGroup.GET("", roomHandlers.ListAllRooms)
	roomGroup.GET("/image/:id", roomHandlers.ListAllRoomsWithImages)
	roomGroup.POST("/image", roomHandlers.AddImageToRoomById)
	roomGroup.GET("/:id", roomHandlers.GetRoomByID)
	roomGroup.GET("/cod/:cod", roomHandlers.GetRoomByCod)
	roomGroup.DELETE("/:id", roomHandlers.DeleteRoomByID)
}
