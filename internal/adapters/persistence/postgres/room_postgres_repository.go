package postgres

import (
	"context"
	"fmt"
	"reservify/internal/adapters/persistence/postgres/bridge"
	"reservify/internal/app/entity/room"
	"reservify/internal/app/interfaces/repository"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var _ repository.RoomLoader = &RoomPostgresRepository{}

type RoomPostgresRepository struct {
	connectorManager
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func (instance RoomPostgresRepository) CreateRoom(u room.Room) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = queries.CreateRoom(ctx, bridge.CreateRoomParams{
		ID:        u.ID(),
		Cod:       u.Cod(),
		Number:    int32(u.Number()),
		Vacancies: int32(u.Vacancies()),
		Price:     FloatToString(u.Price()),
	})

	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %v", err)
	}

	return nil
}

func (instance RoomPostgresRepository) ListAllRooms() ([]room.Room, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	roomsDB, err := queries.ListAllRooms(ctx)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter usuário: %v", err)
	}

	var rooms []room.Room

	for _, roomDB := range roomsDB {
		price, err := StringToFloat(roomDB.Price)

		if err != nil {
			return nil, fmt.Errorf("falha ao obter usuário: %v", err)
		}

		roomBuild, err := room.NewBuilder().WithID(roomDB.ID).WithCod(roomDB.Cod).WithNumber(int(roomDB.Number)).WithVacancies(int(roomDB.Vacancies)).WithPrice(price).Build()

		if err != nil {
			return nil, fmt.Errorf("falha ao obter usuário: %v", err)
		}

		rooms = append(rooms, *roomBuild)
	}

	return rooms, nil
}

func (instance RoomPostgresRepository) GetRoomByID(id uuid.UUID) (*room.Room, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	stmt, err := conn.Prepare("SELECT * FROM room WHERE id = $1;")

	if err != nil {
		return nil, fmt.Errorf("falha ao obter usuário: %v", err)
	}

	var idDB uuid.UUID
	var cod string
	var number int
	var vacancies int
	var price float64
	var createdAt time.Time
	var updatedAt time.Time

	err = stmt.QueryRow(id).Scan(&idDB, &cod, &number, &vacancies, &price, &createdAt, &updatedAt)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter usuário: %v", err)
	}

	roomDB, err := room.NewBuilder().WithID(idDB).WithCod(cod).WithNumber(number).WithVacancies(vacancies).WithPrice(price).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).Build()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter usuário: %v", err)
	}

	return roomDB, nil

}

func (instance RoomPostgresRepository) GetRoomByCod(cod string) (*room.Room, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	roomDB, err := queries.FindRoomByCod(ctx, cod)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter quarto: %v", err)
	}

	price, err := StringToFloat(roomDB.Price)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter quarto: %v", err)
	}

	roomBuild, err := room.NewBuilder().WithID(roomDB.ID).WithCod(roomDB.Cod).WithNumber(int(roomDB.Number)).WithVacancies(int(roomDB.Vacancies)).WithPrice(price).Build()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter quarto: %v", err)
	}

	return roomBuild, nil

}

func (instance RoomPostgresRepository) DeleteRoomByID(id uuid.UUID) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	stmt, err := conn.Prepare("DELETE FROM room WHERE id = $1;")

	if err != nil {
		return fmt.Errorf("falha ao obter usuário: %v", err)
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return fmt.Errorf("falha ao obter usuário: %v", err)
	}

	return nil
}

func NewRoomPostgresRepository(connectorManager connectorManager) *RoomPostgresRepository {
	return &RoomPostgresRepository{
		connectorManager: connectorManager,
	}
}
