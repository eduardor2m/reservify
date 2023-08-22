package response

import (
	"reservify/internal/app/entity/room"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `json:"id"`
	Cod       string    `json:"cod"`
	Number    int       `json:"number"`
	Vacancies int       `json:"vacancies"`
	Price     float64   `json:"price"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Images    *[]Image  `json:"images,omitempty"`
}

func NewRoom(room room.Room) *Room {
	return &Room{
		ID:        room.ID(),
		Cod:       room.Cod(),
		Number:    room.Number(),
		Vacancies: room.Vacancies(),
		Price:     room.Price(),
		CreatedAt: room.CreatedAt().String(),
		UpdatedAt: room.UpdatedAt().String(),
		Images:    NewImage(room.Image),
	}
}
