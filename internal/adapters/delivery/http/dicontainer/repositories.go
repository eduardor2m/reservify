package dicontainer

import (
	"reservify/internal/adapters/persistence/postgres"
	"reservify/internal/app/interfaces/repository"
)

func GetUserRepository() repository.UserLoader {
	return postgres.NewUserPostgresRepository(GetPsqlConnectionManager())
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectorManager {
	return &postgres.DatabaseConnectorManager{}
}
