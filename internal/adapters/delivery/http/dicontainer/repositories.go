package dicontainer

import (
	"reservify/internal/adapters/persistence/postgres"
	"reservify/internal/app/interfaces/repository"
)

func GetUserRepository() repository.UserLoader {
	return postgres.NewUserPostgresRepository(GetPsqlConnectionManager())
}

func GetRoomRepository() repository.RoomLoader {
	return postgres.NewRoomPostgresRepository(GetPsqlConnectionManager())
}

func GetReservationRepository() repository.ReservationLoader {
	return postgres.NewReservationPostgresRepository(GetPsqlConnectionManager())
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectorManager {
	return &postgres.DatabaseConnectorManager{}
}
