package primary

import (
	"reservify/internal/app/entity/room"
)

type RoomManager interface {
	CreateRoom(room room.Room) error
}
