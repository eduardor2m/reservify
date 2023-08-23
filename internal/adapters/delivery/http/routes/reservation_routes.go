package routes

import (
	"reservify/internal/adapters/delivery/http/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadReservationRoutes(group *echo.Group) {
	reservationGroup := group.Group("/reservation")
	reservationHandlers := dicontainer.GetReservationHandler()

	reservationGroup.POST("/admin", reservationHandlers.CreateReservation)
	reservationGroup.POST("", reservationHandlers.CreateMyReservation)
	reservationGroup.GET("", reservationHandlers.ListAllReservations)
	reservationGroup.GET("/:id", reservationHandlers.GetReservationByID)
	reservationGroup.GET("/room/:id_room", reservationHandlers.GetReservationsByRoomID)
	reservationGroup.GET("/user/:id_user", reservationHandlers.GetReservationsByUserID)
	reservationGroup.DELETE("/admin/:id", reservationHandlers.DeleteReservationByID)
	reservationGroup.DELETE("/:id", reservationHandlers.DeleteMyReservationByID)
}
