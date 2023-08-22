package postgres

import (
	"context"
	"fmt"
	"reservify/internal/adapters/persistence/postgres/bridge"
	"reservify/internal/app/entity/reservation"
	"reservify/internal/app/interfaces/repository"
	"reservify/internal/utils/converters"

	"github.com/google/uuid"
)

var _ repository.ReservationLoader = &ReservationPostgresRepository{}

type ReservationPostgresRepository struct {
	connectorManager
}

func checkIfRoomIsAvailable(ctx context.Context, queries bridge.Queries, reservation reservation.Reservation) error {
	reservationsDB, err := queries.ListReservationsByRoomID(ctx, reservation.IDRoom())

	if err != nil {
		return fmt.Errorf("falha ao obter reservas do banco de dados: %v", err)
	}

	newCheckIn, err := converters.ConverterFromStringToTime(reservation.CheckIn())
	if err != nil {
		return fmt.Errorf("falha ao converter data de check-in: %v", err)
	}

	newCheckOut, err := converters.ConverterFromStringToTime(reservation.CheckOut())
	if err != nil {
		return fmt.Errorf("falha ao converter data de check-out: %v", err)
	}

	for _, reservationDB := range reservationsDB {
		dbCheckIn, err := converters.ConverterFromStringToTime(reservationDB.CheckIn)
		if err != nil {
			return fmt.Errorf("falha ao converter data de check-in do banco de dados: %v", err)
		}

		dbCheckOut, err := converters.ConverterFromStringToTime(reservationDB.CheckOut)
		if err != nil {
			return fmt.Errorf("falha ao converter data de check-out do banco de dados: %v", err)
		}

		if (newCheckIn.Equal(dbCheckIn) || newCheckIn.Equal(dbCheckOut)) ||
			(newCheckOut.Equal(dbCheckIn) || newCheckOut.Equal(dbCheckOut)) {
			return fmt.Errorf("falha ao criar reserva: quarto indisponível")
		}
	}

	return nil
}

func (instance ReservationPostgresRepository) CreateReservation(
	reservation reservation.Reservation,
) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = checkIfRoomIsAvailable(ctx, *queries, reservation)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = queries.CreateReservation(ctx, bridge.CreateReservationParams{
		ID:       reservation.ID(),
		IDUser:   reservation.IDUser(),
		IDRoom:   reservation.IDRoom(),
		CheckIn:  reservation.CheckIn(),
		CheckOut: reservation.CheckOut(),
	})

	if err != nil {
		return fmt.Errorf("falha ao criar reserva: %v", err)
	}

	return nil
}

func (instance ReservationPostgresRepository) ListAllReservations() ([]reservation.Reservation, error) {
	var reservations []reservation.Reservation

	conn, err := instance.getConnection()

	if err != nil {
		return reservations, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	reservationsDB, err := queries.ListAllReservations(ctx)

	if err != nil {
		return reservations, fmt.Errorf("falha ao listar reservas: %v", err)
	}

	for _, reservationDB := range reservationsDB {
		reservationReceived, err := reservation.NewBuilder().
			WithID(reservationDB.ID).
			WithIdUser(reservationDB.IDUser).
			WithIdRoom(reservationDB.IDRoom).
			WithCheckIn(reservationDB.CheckIn).
			WithCheckOut(reservationDB.CheckOut).
			WithCreatedAt(reservationDB.CreatedAt).
			WithUpdatedAt(reservationDB.UpdatedAt).
			Build()

		if err != nil {
			return reservations, fmt.Errorf("falha ao construir reserva: %v", err)
		}

		reservations = append(reservations, *reservationReceived)
	}

	return reservations, nil
}

func (instance ReservationPostgresRepository) GetReservationByID(id uuid.UUID) (*reservation.Reservation, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	reservationDB, err := queries.GetReservationByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter reserva: %v", err)
	}

	reservationReceived, err := reservation.NewBuilder().
		WithID(reservationDB.ID).
		WithIdUser(reservationDB.IDUser).
		WithIdRoom(reservationDB.IDRoom).
		WithCheckIn(reservationDB.CheckIn).
		WithCheckOut(reservationDB.CheckOut).
		Build()

	if err != nil {
		return nil, fmt.Errorf("falha ao construir reserva: %v", err)
	}

	return reservationReceived, nil
}

func (instance ReservationPostgresRepository) GetReservationsByRoomID(roomID uuid.UUID) ([]reservation.Reservation, error) {
	var reservations []reservation.Reservation

	conn, err := instance.getConnection()

	if err != nil {
		return reservations, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	reservationsDB, err := queries.ListReservationsByRoomID(ctx, roomID)

	if err != nil {
		return reservations, fmt.Errorf("falha ao obter reservas: %v", err)
	}

	for _, reservationDB := range reservationsDB {
		reservationReceived, err := reservation.NewBuilder().
			WithID(reservationDB.ID).
			WithIdUser(reservationDB.IDUser).
			WithIdRoom(reservationDB.IDRoom).
			WithCheckIn(reservationDB.CheckIn).
			WithCheckOut(reservationDB.CheckOut).
			Build()

		if err != nil {
			return reservations, fmt.Errorf("falha ao construir reserva: %v", err)
		}

		reservations = append(reservations, *reservationReceived)
	}

	return reservations, nil
}

func (instance ReservationPostgresRepository) GetReservationsByUserID(userID uuid.UUID) ([]reservation.Reservation, error) {
	var reservations []reservation.Reservation

	conn, err := instance.getConnection()

	if err != nil {
		return reservations, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	reservationsDB, err := queries.ListReservationsByUserID(ctx, userID)

	if err != nil {
		return reservations, fmt.Errorf("falha ao obter reservas: %v", err)
	}

	for _, reservationDB := range reservationsDB {
		reservationReceived, err := reservation.NewBuilder().
			WithID(reservationDB.ID).
			WithIdUser(reservationDB.IDUser).
			WithIdRoom(reservationDB.IDRoom).
			WithCheckIn(reservationDB.CheckIn).
			WithCheckOut(reservationDB.CheckOut).
			Build()

		if err != nil {
			return reservations, fmt.Errorf("falha ao construir reserva: %v", err)
		}

		reservations = append(reservations, *reservationReceived)
	}

	return reservations, nil
}

func (instance ReservationPostgresRepository) DeleteReservationByID(id uuid.UUID) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = queries.DeleteReservation(ctx, id)

	if err != nil {
		return fmt.Errorf("falha ao deletar reserva: %v", err)
	}

	return nil
}

func NewReservationPostgresRepository(connectorManager connectorManager) *ReservationPostgresRepository {
	return &ReservationPostgresRepository{
		connectorManager: connectorManager,
	}
}
