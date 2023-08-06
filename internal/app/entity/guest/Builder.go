package guest

import (
	"github.com/google/uuid"
	"time"
)

type Builder struct {
	Guest Guest
	Err   error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithID(id uuid.UUID) *Builder {
	instance.Guest.id = id
	return instance
}

func (instance *Builder) WithCPF(CPF string) *Builder {
	instance.Guest.cpf = CPF
	return instance
}

func (instance *Builder) WithName(name string) *Builder {
	instance.Guest.name = name
	return instance
}

func (instance *Builder) WithCreatedAt(createdAt time.Time) *Builder {
	instance.Guest.createdAt = createdAt
	return instance
}

func (instance *Builder) WithUpdatedAt(updatedAt time.Time) *Builder {
	instance.Guest.updatedAt = updatedAt
	return instance
}

func (instance *Builder) Build() (*Guest, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}
	return &instance.Guest, nil
}
