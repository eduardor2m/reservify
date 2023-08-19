package reservation

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	id        uuid.UUID
	idUser    uuid.UUID
	idRoom    uuid.UUID
	checkIn   string
	checkOut  string
	createdAt time.Time
	updatedAt time.Time
}

func (instance *Reservation) ID() uuid.UUID {
	return instance.id
}

func (instance *Reservation) IDUser() uuid.UUID {
	return instance.idUser
}

func (instance *Reservation) IDRoom() uuid.UUID {
	return instance.idRoom
}

func (instance *Reservation) CheckIn() string {
	return instance.checkIn
}

func (instance *Reservation) CheckOut() string {
	return instance.checkOut
}

func (instance *Reservation) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Reservation) UpdatedAt() time.Time {
	return instance.updatedAt
}
