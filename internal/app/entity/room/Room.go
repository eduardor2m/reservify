package room

import (
	"reservify/internal/app/entity/image"
	"time"

	"github.com/google/uuid"
)

type Room struct {
	id          uuid.UUID
	name        string
	description string
	cod         string
	number      int
	vacancies   int
	price       float64
	createdAt   time.Time
	updatedAt   time.Time
	Image       []image.Image
}

func (instance *Room) ID() uuid.UUID {
	if instance.id == uuid.Nil {
		instance.id = uuid.New()
	}

	return instance.id
}

func (instance *Room) Name() string {
	return instance.name
}

func (instance *Room) Description() string {
	return instance.description
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
