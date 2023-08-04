package user

import "github.com/google/uuid"

type Builder struct {
	User User
	Err  error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithID(id uuid.UUID) *Builder {
	instance.User.id = id
	return instance
}

func (instance *Builder) WithName(name string) *Builder {
	instance.User.name = name
	return instance
}

func (instance *Builder) WithEmail(email string) *Builder {
	instance.User.email = email
	return instance
}

func (instance *Builder) WithPassword(password string) *Builder {
	instance.User.password = password
	return instance
}

func (instance *Builder) WithDateOfBirth(dateOfBirth string) *Builder {
	instance.User.dateOfBirth = dateOfBirth
	return instance
}

func (instance *Builder) WithAdmin(admin bool) *Builder {
	instance.User.admin = admin
	return instance
}

func (instance *Builder) Build() (*User, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}
	return &instance.User, nil
}
