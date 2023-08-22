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

	return instance.reservationRepository.CreateReservation(*reservationFormatted)
}

func (instance *ReservationServices) ListAllReservations() ([]reservation.Reservation, error) {
	return instance.reservationRepository.ListAllReservations()
}

func (instance *ReservationServices) GetReservationByID(id uuid.UUID) (*reservation.Reservation, error) {
	return instance.reservationRepository.GetReservationByID(id)
}

func (instance *ReservationServices) GetReservationsByRoomID(roomID uuid.UUID) ([]reservation.Reservation, error) {
	return instance.reservationRepository.GetReservationsByRoomID(roomID)
}

func (instance *ReservationServices) GetReservationsByUserID(userID uuid.UUID) ([]reservation.Reservation, error) {
	return instance.reservationRepository.GetReservationsByUserID(userID)
}

func (instance *ReservationServices) DeleteReservationByID(id uuid.UUID) error {
	return instance.reservationRepository.DeleteReservationByID(id)
}

func NewReservationServices(
	reservationRepository repository.ReservationLoader,
) *ReservationServices {
	return &ReservationServices{
		reservationRepository: reservationRepository,
	}
}
