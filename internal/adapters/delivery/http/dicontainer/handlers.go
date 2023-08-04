package dicontainer

import "reservify/internal/adapters/delivery/http/handlers"

func GetUserHandler() *handlers.UserHandler {
	return handlers.NewUserHandler(GetUserServices())
}
