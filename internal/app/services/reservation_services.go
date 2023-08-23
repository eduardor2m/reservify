package services

import (
	"reservify/internal/app/entity/reservation"
	"reservify/internal/app/interfaces/primary"
	"reservify/internal/app/interfaces/repository"
	"time"

	"github.com/google/uuid"
)

var _ primary.ReservationManager = (*ReservationServices)(nil)

type ReservationServices struct {
	reservationRepository repository.ReservationLoader
}

func (instance *ReservationServices) CreateReservation(
	r reservation.Reservation,
	tokenJwt string,
) error {
	reservationUUID, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	reservationTime := time.Now()

	reservationFormatted, err := reservation.NewBuilder().WithID(reservationUUID).WithIdRoom(r.IDRoom()).WithIdUser(r.IDUser()).WithCheckIn(r.CheckIn()).WithCheckOut(r.CheckOut()).WithCreatedAt(reservationTime).WithUpdatedAt(reservationTime).Build()

	if err != nil {
		return err
	}

	return instance.reservationRepository.CreateReservation(*reservationFormatted, tokenJwt)
}

func (instance *ReservationServices) CreateMyReservation(
	r reservation.Reservation,
	tokenJwt string,
) error {
	reservationUUID, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	reservationTime := time.Now()

	reservationFormatted, err := reservation.NewBuilder().WithID(reservationUUID).WithIdRoom(r.IDRoom()).WithIdUser(r.IDUser()).WithCheckIn(r.CheckIn()).WithCheckOut(r.CheckOut()).WithCreatedAt(reservationTime).WithUpdatedAt(reservationTime).Build()

	if err != nil {
		return err
	}

	return instance.reservationRepository.CreateMyReservation(*reservationFormatted, tokenJwt)
}

func (instance *ReservationServices) ListAllReservations(tokenJwt string) ([]reservation.Reservation, error) {
	return instance.reservationRepository.ListAllReservations(tokenJwt)
}

func (instance *ReservationServices) GetReservationByID(id uuid.UUID, tokenJwt string) (*reservation.Reservation, error) {
	return instance.reservationRepository.GetReservationByID(id, tokenJwt)
}

func (instance *ReservationServices) GetReservationsByRoomID(roomID uuid.UUID, tokenJwt string) ([]reservation.Reservation, error) {
	return instance.reservationRepository.GetReservationsByRoomID(roomID, tokenJwt)
}

func (instance *ReservationServices) GetReservationsByUserID(userID uuid.UUID) ([]reservation.Reservation, error) {
	return instance.reservationRepository.GetReservationsByUserID(userID)
}

func (instance *ReservationServices) DeleteReservationByID(id uuid.UUID, tokenJwt string) error {
	return instance.reservationRepository.DeleteReservationByID(id, tokenJwt)
}

func (instance *ReservationServices) DeleteMyReservationByID(id uuid.UUID, tokenJwt string) error {
	return instance.reservationRepository.DeleteMyReservationByID(id, tokenJwt)
}

func NewReservationServices(
	reservationRepository repository.ReservationLoader,
) *ReservationServices {
	return &ReservationServices{
		reservationRepository: reservationRepository,
	}
}
