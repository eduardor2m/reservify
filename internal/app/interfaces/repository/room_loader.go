package repository

import "reservify/internal/app/entity/room"

type RoomLoader interface {
	CreateRoom(room room.Room) error
}
