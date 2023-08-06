package room

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	id        uuid.UUID
	cod       string
	number    int
	createdAt time.Time
	updatedAt time.Time
}

func (instance *Room) ID() uuid.UUID {
	return instance.id
}

func (instance *Room) Cod() string {
	return instance.cod
}

func (instance *Room) Number() int {
	return instance.number
}

func (instance *Room) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Room) UpdatedAt() time.Time {
	return instance.updatedAt
}
