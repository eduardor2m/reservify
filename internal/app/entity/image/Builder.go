package image

import (
	"github.com/google/uuid"
)

type Builder struct {
	Image Image
	Err   error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithIDRoom(idRoom uuid.UUID) *Builder {
	instance.Image.idRoom = idRoom
	return instance
}

func (instance *Builder) WithImageUrl(imageUrl string) *Builder {
	instance.Image.imageUrl = imageUrl
	return instance
}

func (instance *Builder) WithThumbnail(thumbnail bool) *Builder {
	instance.Image.thumbnail = thumbnail
	return instance
}

func (instance *Builder) Build() (*Image, error) {
	if instance.Err != nil {
		return nil, instance.Err
	}
	return &instance.Image, nil
}
