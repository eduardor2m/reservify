package guest

import (
	"github.com/google/uuid"
	"time"
)

type Guest struct {
	id        uuid.UUID
	cpf       string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func (instance *Guest) ID() uuid.UUID {
	return instance.id
}

func (instance *Guest) CPF() string {
	return instance.cpf

}

func (instance *Guest) Name() string {
	return instance.name
}

func (instance *Guest) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Guest) UpdatedAt() time.Time {
	return instance.updatedAt
}
