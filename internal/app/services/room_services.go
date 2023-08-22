package services

import (
	"reservify/internal/app/entity/room"
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/interfaces/repository"
	"time"

	"github.com/google/uuid"
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

	createAt := time.Now()

	formattedRoom, err := room.NewBuilder().WithID(newRoomUUID).WithCod(r.Cod()).WithNumber(r.Number()).WithVacancies(r.Vacancies()).WithPrice(r.Price()).WithCreatedAt(createAt).WithUpdatedAt(createAt).Build()

	if err != nil {
		return err
	}

	return instance.roomRepository.CreateRoom(*formattedRoom)
}

func (instance *RoomServices) ListAllRooms() ([]room.Room, error) {
	return instance.roomRepository.ListAllRooms()
}

func (instance *RoomServices) GetRoomByID(id uuid.UUID) (*room.Room, error) {
	return instance.roomRepository.GetRoomByID(id)
}

func (instance *RoomServices) GetRoomByCod(cod string) (*room.Room, error) {
	return instance.roomRepository.GetRoomByCod(cod)
}

func (instance *RoomServices) DeleteRoomByID(id uuid.UUID) error {
	return instance.roomRepository.DeleteRoomByID(id)
}

func (instance *RoomServices) AddImageToRoomById(id uuid.UUID, image_url string) error {
	return instance.roomRepository.AddImageToRoomById(id, image_url)
}

func (instance *RoomServices) ListAllImagesByRoomID(id uuid.UUID) (*room.Room, error) {
	return instance.roomRepository.ListAllImagesByRoomID(id)
}

func NewRoomServices(roomRepository repository.RoomLoader) *RoomServices {
	return &RoomServices{
		roomRepository: roomRepository,
	}
}
