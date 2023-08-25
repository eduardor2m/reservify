package primary

import (
	"reservify/internal/app/entity/room"

	"github.com/google/uuid"
)

type RoomManager interface {
	CreateRoom(room room.Room, tokenJwt string) error
	ListAllRooms() ([]room.Room, error)
	AddImageToRoomByRoomID(id uuid.UUID, imageUrl string, imageThumbnail bool, tokenJwt string) error
	GetRoomWithImagesByRoomID(id uuid.UUID) (*room.Room, error)
	GetRoomByID(id uuid.UUID) (*room.Room, error)
	GetRoomByCod(cod string) (*room.Room, error)
	DeleteRoomByID(id uuid.UUID, tokenJwt string) error
}
