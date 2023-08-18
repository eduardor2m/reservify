package reservation

import (
	"github.com/google/uuid"
	"time"
)

type Guest struct {
	id        uuid.UUID
	idUser    uuid.UUID
	idRoom    uuid.UUID
	checkIn   time.Time
	checkOut  time.Time
	createdAt time.Time
	updatedAt time.Time
}

func (instance *Guest) ID() uuid.UUID {
	return instance.id
}

func (instance *Guest) IDUser() uuid.UUID {
	return instance.idUser
}

func (instance *Guest) IDRoom() uuid.UUID {
	return instance.idRoom
}

func (instance *Guest) CheckIn() time.Time {
	return instance.checkIn
}

func (instance *Guest) CheckOut() time.Time {
	return instance.checkOut
}

func (instance *Guest) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Guest) UpdatedAt() time.Time {
	return instance.updatedAt
}
