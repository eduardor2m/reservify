package image

import (
	"github.com/google/uuid"
)

type Image struct {
	idRoom    uuid.UUID
	imageUrl	string
}

func (instance *Image) IDRoom() uuid.UUID {
	return instance.idRoom
}

func (instance *Image) ImageUrl() string {
	return instance.imageUrl
}

