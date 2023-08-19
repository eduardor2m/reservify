package user

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	User User
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

	instance.User.id = id
	return instance
}

func (instance *Builder) WithCPF(cpf string) *Builder {
	if len(cpf) < 11 {
		instance.Err = fmt.Errorf("cpf deve conter no mínimo 11 caracteres")
		return instance
	}

	instance.User.cpf = cpf
	return instance
}

func (instance *Builder) WithName(name string) *Builder {
	if len(name) < 3 {
		instance.Err = fmt.Errorf("nome deve conter no mínimo 3 caracteres")
		return instance
	}

	instance.User.name = name
	return instance
}

func (instance *Builder) WithEmail(email string) *Builder {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(regex, email)

	if err != nil {
		instance.Err = fmt.Errorf("falha ao validar email: %v", err)
		return instance
	}

	if !match {
		instance.Err = fmt.Errorf("email inválido")
		return instance
	}

	if len(email) < 8 {
		instance.Err = fmt.Errorf("email deve conter no mínimo 8 caracteres")
		return instance
	}

	instance.User.email = email
	return instance
}

func (instance *Builder) WithPhone(phone string) *Builder {
	if len(phone) < 8 {
		instance.Err = fmt.Errorf("telefone deve conter no mínimo 8 caracteres")
		return instance
	}

	instance.User.phone = phone
	return instance
}

func (instance *Builder) WithPassword(password string) *Builder {
	if len(password) < 8 {
		instance.Err = fmt.Errorf("senha deve conter no mínimo 8 caracteres")
		return instance
	}

	instance.User.password = password
	return instance
}

func (instance *Builder) WithDateOfBirth(dateOfBirth string) *Builder {
	if len(dateOfBirth) < 10 {
		instance.Err = fmt.Errorf("data de nascimento deve conter no mínimo 10 caracteres")
	}

	regex := `^(\d{2})\/(\d{2})\/(\d{4})$`

	match, err := regexp.MatchString(regex, dateOfBirth)

	if err != nil {
		instance.Err = fmt.Errorf("falha ao validar data de nascimento: %v", err)
		return instance
	}

	if !match {
		instance.Err = fmt.Errorf("data de nascimento inválida")
		return instance
	}

	instance.User.dateOfBirth = dateOfBirth
	return instance
}

func (instance *Builder) WithAdmin(admin bool) *Builder {
	instance.User.admin = admin
	return instance
}

func (instance *Builder) WithCreatedAt(createdAt time.Time) *Builder {
	instance.User.createdAt = createdAt
	return instance
}

func (instance *Builder) WithUpdatedAt(updatedAt time.Time) *Builder {
	instance.User.updatedAt = updatedAt
	return instance
}

func (instance *Builder) Build() (*User, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}
	return &instance.User, nil
}
