package room

import (
	"github.com/google/uuid"
	"time"
)

type Builder struct {
	Room Room
	Err  error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithID(id uuid.UUID) *Builder {
	instance.Room.id = id
	return instance
}

func (instance *Builder) WithCod(cod string) *Builder {
	instance.Room.cod = cod
	return instance
}

func (instance *Builder) WithNumber(number int) *Builder {
	instance.Room.number = number
	return instance
}

func (instance *Builder) WithVacancies(vacancies int) *Builder {
	instance.Room.vacancies = vacancies
	return instance
}

func (instance *Builder) WithPrice(price float64) *Builder {
	instance.Room.price = price
	return instance
}

func (instance *Builder) WithCreatedAt(createdAt time.Time) *Builder {
	instance.Room.createdAt = createdAt
	return instance
}

func (instance *Builder) WithUpdatedAt(updatedAt time.Time) *Builder {
	instance.Room.updatedAt = updatedAt
	return instance
}

func (instance *Builder) Build() (*Room, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}
	return &instance.Room, nil
}
