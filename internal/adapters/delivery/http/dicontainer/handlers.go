package dicontainer

import "reservify/internal/adapters/delivery/http/handlers"

func GetUserHandler() *handlers.UserHandler {
	return handlers.NewUserHandler(GetUserServices())
}

func GetRoomHandler() *handlers.RoomHandler {
	return handlers.NewRoomHandler(GetRoomServices())
}

func GetReservationHandler() *handlers.ReservationHandler {
	return handlers.NewReservationHandler(GetReservationServices())
}
