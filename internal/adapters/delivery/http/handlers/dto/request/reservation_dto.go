package request

import "github.com/google/uuid"

type ReservationDTO struct {
	IdUser   uuid.UUID `json:"id_user"`
	IdRoom   uuid.UUID `json:"id_room"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}
