package services

import (
	"github.com/google/uuid"
	"reservify/internal/app/entity/room"
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/interfaces/repository"
)

var _ primary.RoomManager = (*RoomServices)(nil)

type RoomServices struct {
	roomRepository repository.RoomLoader
}

func (instance *RoomServices) CreateRoom(r room.Room) error {
	newRoomUUID, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	formattedRoom, err := room.NewBuilder().WithID(newRoomUUID).WithCod(r.Cod()).WithNumber(r.Number()).WithVacancies(r.Vacancies()).WithPrice(r.Price()).Build()

	return instance.roomRepository.CreateRoom(*formattedRoom)
}

func (instance *RoomServices) ListAllRooms() ([]room.Room, error) {
	return instance.roomRepository.ListAllRooms()
}

func (instance *RoomServices) GetRoomByID(id uuid.UUID) (*room.Room, error) {
	return instance.roomRepository.GetRoomByID(id)
}

func NewRoomServices(roomRepository repository.RoomLoader) *RoomServices {
	return &RoomServices{
		roomRepository: roomRepository,
	}
}
