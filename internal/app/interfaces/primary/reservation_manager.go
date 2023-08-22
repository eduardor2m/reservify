package primary

import (
	"reservify/internal/app/entity/reservation"

	"github.com/google/uuid"
)

type ReservationManager interface {
	CreateReservation(reservation reservation.Reservation) error
	ListAllReservations() ([]reservation.Reservation, error)
	GetReservationByID(id uuid.UUID) (*reservation.Reservation, error)
	GetReservationsByRoomID(roomID uuid.UUID) ([]reservation.Reservation, error)
	GetReservationsByUserID(userID uuid.UUID) ([]reservation.Reservation, error)
	DeleteReservationByID(id uuid.UUID) error
}
