package postgres

import (
	"context"
	"fmt"
	"os"
	"reservify/internal/adapters/persistence/postgres/bridge"
	"reservify/internal/app/entity/reservation"
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/repository"
	"reservify/internal/utils/converters"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var _ repository.UserLoader = &UserPostgresRepository{}

type UserPostgresRepository struct {
	connectorManager
}

func checkIfRoomIsAvailable(ctx context.Context, queries bridge.Queries, reservation reservation.Reservation) error {
	reservationsDB, err := queries.GetReservationByIDRoom(ctx, reservation.IDRoom())

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

func (instance UserPostgresRepository) CreateUser(u user.User) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	err = queries.CreateUser(ctx, bridge.CreateUserParams{
		ID:          u.ID(),
		Name:        u.Name(),
		Cpf:         u.CPF(),
		Email:       u.Email(),
		Password:    u.Password(),
		Phone:       u.Phone(),
		DateOfBirth: u.DateOfBirth(),
		Admin:       u.Admin(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %v", err)
	}

	return nil
}

func (instance UserPostgresRepository) LoginUser(email string, password string) (*string, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}
	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	userDB, err := queries.Login(ctx, email)

	if err != nil {
		return nil, fmt.Errorf("falha ao logar usuário: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("falha ao comparar senha: %v", err)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return nil, fmt.Errorf("falha ao criar token: %v", err)
	}

	return &tokenString, nil
}

func (instance UserPostgresRepository) CreateReservation(
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

func (instance UserPostgresRepository) GetReservationByID(id uuid.UUID) (*reservation.Reservation, error) {
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

func (instance UserPostgresRepository) GetReservationByIDRoom(id uuid.UUID) ([]reservation.Reservation, error) {
	var reservations []reservation.Reservation

	conn, err := instance.getConnection()

	if err != nil {
		return reservations, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	reservationsDB, err := queries.GetReservationByIDRoom(ctx, id)

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

func (instance UserPostgresRepository) GetReservationByIDUser(id uuid.UUID) ([]reservation.Reservation, error) {
	var reservations []reservation.Reservation

	conn, err := instance.getConnection()

	if err != nil {
		return reservations, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	reservationsDB, err := queries.GetReservationByIDUser(ctx, id)

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

func (instance UserPostgresRepository) DeleteReservationByID(id uuid.UUID) error {
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

func (instance UserPostgresRepository) ListAllReservations() ([]reservation.Reservation, error) {
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

func (instance UserPostgresRepository) ListAllUsers() ([]user.User, error) {
	var users []user.User

	conn, err := instance.getConnection()

	if err != nil {
		return users, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	usersDB, err := queries.ListAll(ctx)

	if err != nil {
		return users, fmt.Errorf("falha ao listar usuários: %v", err)
	}

	for _, userDB := range usersDB {
		userReceived, err := user.NewBuilder().
			WithID(userDB.ID).
			WithName(userDB.Name).
			WithEmail(userDB.Email).
			WithPassword(userDB.Password).
			WithDateOfBirth(userDB.DateOfBirth).
			WithCPF(userDB.Cpf).
			WithPhone(userDB.Phone).
			WithAdmin(userDB.Admin).
			WithCreatedAt(userDB.CreatedAt).
			WithUpdatedAt(userDB.UpdatedAt).
			Build()

		if err != nil {
			return users, fmt.Errorf("falha ao construir usuário: %v", err)
		}

		users = append(users, *userReceived)
	}

	return users, nil
}

func (instance UserPostgresRepository) GetUserByID(id uuid.UUID) (*user.User, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	userDB, err := queries.FindByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter usuário: %v", err)
	}

	userReceived, err := user.NewBuilder().
		WithID(userDB.ID).
		WithName(userDB.Name).
		WithEmail(userDB.Email).
		WithPassword(userDB.Password).
		WithDateOfBirth(userDB.DateOfBirth).
		WithCPF(userDB.Cpf).
		WithPhone(userDB.Phone).
		WithAdmin(userDB.Admin).
		WithCreatedAt(userDB.CreatedAt).
		WithUpdatedAt(userDB.UpdatedAt).
		Build()

	if err != nil {
		return nil, fmt.Errorf("falha ao construir usuário: %v", err)
	}

	return userReceived, nil
}

func (instance UserPostgresRepository) GetUserByName(name string) ([]user.User, error) {
	var users []user.User

	conn, err := instance.getConnection()

	if err != nil {
		return users, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	usersDB, err := queries.ListByName(ctx, name)

	if err != nil {
		return users, fmt.Errorf("falha ao listar usuários: %v", err)
	}

	for _, userDB := range usersDB {
		userReceived, err := user.NewBuilder().
			WithID(userDB.ID).
			WithName(userDB.Name).
			WithEmail(userDB.Email).
			WithPassword(userDB.Password).
			WithDateOfBirth(userDB.DateOfBirth).
			WithCPF(userDB.Cpf).
			WithPhone(userDB.Phone).
			WithAdmin(userDB.Admin).
			WithCreatedAt(userDB.CreatedAt).
			WithUpdatedAt(userDB.UpdatedAt).
			Build()

		if err != nil {
			return users, fmt.Errorf("falha ao construir usuário: %v", err)
		}

		users = append(users, *userReceived)
	}

	return users, nil
}

func (instance UserPostgresRepository) UpdateUserByEmail(email string, user user.User) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = queries.UpdateByEmail(ctx, bridge.UpdateByEmailParams{
		Name:        user.Name(),
		Cpf:         user.CPF(),
		Email:       user.Email(),
		Password:    user.Password(),
		Phone:       user.Phone(),
		DateOfBirth: user.DateOfBirth(),
		Admin:       user.Admin(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return fmt.Errorf("falha ao atualizar usuário: %v", err)
	}

	return nil
}

func (instance UserPostgresRepository) DeleteUserByEmail(email string) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = queries.DeleteByEmail(ctx, email)

	if err != nil {
		return fmt.Errorf("falha ao deletar usuário: %v", err)
	}

	return nil
}

func NewUserPostgresRepository(connectorManager connectorManager) *UserPostgresRepository {
	return &UserPostgresRepository{
		connectorManager: connectorManager,
	}
}
