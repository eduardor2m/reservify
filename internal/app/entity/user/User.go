package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id          uuid.UUID
	name        string
	email       string
	password    string
	dateOfBirth string
	admin       bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *User) ID() uuid.UUID {
	return instance.id
}

func (instance *User) Name() string {
	return instance.name
}

func (instance *User) Email() string {
	return instance.email
}

func (instance *User) Password() string {
	return instance.password
}

func (instance *User) DateOfBirth() string {
	return instance.dateOfBirth
}

func (instance *User) Admin() bool {
	return instance.admin
}

func (instance *User) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *User) UpdatedAt() time.Time {
	return instance.updatedAt
}
