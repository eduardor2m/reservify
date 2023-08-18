package room

import (
	"fmt"
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
	if id == uuid.Nil {
		instance.Err = fmt.Errorf("id inválido")
		return instance
	}

	instance.Room.id = id
	return instance
}

func (instance *Builder) WithCod(cod string) *Builder {
	if len(cod) < 3 {
		instance.Err = fmt.Errorf("código deve conter no mínimo 3 caracteres")
		return instance
	}

	instance.Room.cod = cod
	return instance
}

func (instance *Builder) WithNumber(number int) *Builder {
	if number < 1 {
		instance.Err = fmt.Errorf("número deve ser maior que 0")
		return instance
	}

	instance.Room.number = number
	return instance
}

func (instance *Builder) WithVacancies(vacancies int) *Builder {
	if vacancies < 1 {
		instance.Err = fmt.Errorf("vagas deve ser maior que 0")
		return instance
	}

	instance.Room.vacancies = vacancies
	return instance
}

func (instance *Builder) WithPrice(price float64) *Builder {
	if price < 1 {
		instance.Err = fmt.Errorf("preço deve ser maior que 0")
		return instance
	}

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
