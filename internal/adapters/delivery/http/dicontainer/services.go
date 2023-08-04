package dicontainer

import (
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/services"
)

func GetUserServices() primary.UserManager {
	return services.NewUserServices(GetUserRepository())
}
