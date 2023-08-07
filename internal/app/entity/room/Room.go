package room

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	id        uuid.UUID
	cod       string
	number    int
	vacancies int
	price     float64
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

func (instance *Room) Vacancies() int {
	return instance.vacancies
}

func (instance *Room) Price() float64 {
	return instance.price
}

func (instance *Room) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Room) UpdatedAt() time.Time {
	return instance.updatedAt
}
