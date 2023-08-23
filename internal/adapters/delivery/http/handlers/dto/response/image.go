package response

import (
	"reservify/internal/app/entity/image"

	"github.com/google/uuid"
)

type Image struct {
	IdUser   uuid.UUID `json:"id_room"`
	ImageUrl string    `json:"image_url"`
}

func NewImage(image []image.Image) *[]Image {
	if image == nil {
		return nil
	}

	var images []Image
	for _, image := range image {
		images = append(images, Image{
			IdUser:   image.IDRoom(),
			ImageUrl: image.ImageUrl(),
		})
	}
	return &images
}
