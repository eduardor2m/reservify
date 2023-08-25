package request

import "github.com/google/uuid"

type ImageDTO struct {
	IdUser    uuid.UUID `json:"id_room"`
	ImageUrl  string    `json:"image_url"`
	Thumbnail bool      `json:"thumbnail"`
}
