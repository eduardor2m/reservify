package reservation

import (
	"github.com/google/uuid"
	"time"
)

type Builder struct {
	Reservation Guest
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

func (instance *Builder) WithCreatedAt(createdAt time.Time) *Builder {
	instance.Reservation.createdAt = createdAt
	return instance
}

func (instance *Builder) WithUpdatedAt(updatedAt time.Time) *Builder {
	instance.Reservation.updatedAt = updatedAt
	return instance
}

func (instance *Builder) Build() (*Guest, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}
	return &instance.Reservation, nil
}
