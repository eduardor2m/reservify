package repository

import (
	"github.com/google/uuid"
	"reservify/internal/app/entity/room"
)

type RoomLoader interface {
	CreateRoom(room room.Room) error
	ListAllRooms() ([]room.Room, error)
	GetRoomByID(id uuid.UUID) (*room.Room, error)
}
