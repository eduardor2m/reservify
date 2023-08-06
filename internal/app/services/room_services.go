package services

import (
	"reservify/internal/app/entity/room"
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/interfaces/repository"
)

var _ primary.RoomManager = (*RoomServices)(nil)

type RoomServices struct {
	roomRepository repository.RoomLoader
}

func (instance *RoomServices) CreateRoom(r room.Room) error {
	return instance.roomRepository.CreateRoom(r)
}

func NewRoomServices(roomRepository repository.RoomLoader) *RoomServices {
	return &RoomServices{
		roomRepository: roomRepository,
	}
}
