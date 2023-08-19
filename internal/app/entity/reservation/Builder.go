package reservation

import (
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	Reservation Reservation
	Err         error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithID(id uuid.UUID) *Builder {
	instance.Reservation.id = id
	return instance
}

func (instance *Builder) WithIdRoom(idRoom uuid.UUID) *Builder {
	instance.Reservation.idRoom = idRoom
	return instance
}

func (instance *Builder) WithIdUser(idUser uuid.UUID) *Builder {
	instance.Reservation.idUser = idUser
	return instance
}

func (instance *Builder) WithCheckIn(checkIn string) *Builder {
	instance.Reservation.checkIn = checkIn
	return instance
}

func (instance *Builder) WithCheckOut(checkOut string) *Builder {
	instance.Reservation.checkOut = checkOut
	return instance
}

func (instance *Builder) WithCreatedAt(createdAt time.Time) *Builder {
	instance.Reservation.createdAt = createdAt
	return instance
}

func (instance *Builder) WithUpdatedAt(updatedAt time.Time) *Builder {
	instance.Reservation.updatedAt = updatedAt
	return instance
}

func (instance *Builder) Build() (*Reservation, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}
	return &instance.Reservation, nil
}
