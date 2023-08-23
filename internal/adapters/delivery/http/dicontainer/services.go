package dicontainer

import (
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/services"
)

func GetUserServices() primary.UserManager {
	return services.NewUserServices(GetUserRepository())
}

func GetRoomServices() primary.RoomManager {
	return services.NewRoomServices(GetRoomRepository())
}

func GetReservationServices() primary.ReservationManager {
	return services.NewReservationServices(GetReservationRepository())
}
