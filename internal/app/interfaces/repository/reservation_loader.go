package repository

import (
	"reservify/internal/app/entity/reservation"

	"github.com/google/uuid"
)

type ReservationLoader interface {
	CreateReservation(reservation reservation.Reservation, tokenJwt string) error
	CreateMyReservation(reservation reservation.Reservation, tokenJwt string) error
	ListAllReservations(tokenJwt string) ([]reservation.Reservation, error)
	GetReservationByID(id uuid.UUID, tokenJwt string) (*reservation.Reservation, error)
	GetReservationsByRoomID(roomID uuid.UUID, tokenJwt string) ([]reservation.Reservation, error)
	GetReservationsByUserID(userID uuid.UUID) ([]reservation.Reservation, error)
	DeleteReservationByID(id uuid.UUID, tokenJwt string) error
	DeleteMyReservationByID(id uuid.UUID, tokenJwt string) error
}
