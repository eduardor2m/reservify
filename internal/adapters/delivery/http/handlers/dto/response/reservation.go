package response

import (
	"reservify/internal/app/entity/reservation"

	"github.com/google/uuid"
)

type Reservation struct {
	ID        uuid.UUID `json:"id"`
	IDUser    uuid.UUID `json:"id_user"`
	IDRoom    uuid.UUID `json:"id_room"`
	CheckIn   string    `json:"check_in"`
	CheckOut  string    `json:"check_out"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewReservation(reservation reservation.Reservation) *Reservation {
	return &Reservation{
		ID:        reservation.ID(),
		IDUser:    reservation.IDUser(),
		IDRoom:    reservation.IDRoom(),
		CheckIn:   reservation.CheckIn(),
		CheckOut:  reservation.CheckOut(),
		CreatedAt: reservation.CreatedAt().String(),
		UpdatedAt: reservation.UpdatedAt().String(),
	}
}
