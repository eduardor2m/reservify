package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"reservify/internal/app/entity/room"
	"reservify/internal/app/interfaces/repository"
	"time"
)

var _ repository.RoomLoader = &RoomPostgresRepository{}

type RoomPostgresRepository struct {
	connectorManager
}

func (instance RoomPostgresRepository) CreateRoom(u room.Room) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	stmt, err := conn.Prepare("INSERT INTO room (id, cod, number, vacancies, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);")

	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %v", err)
	}

	_, err = stmt.Exec(u.ID(), u.Cod(), u.Number(), u.Vacancies(), u.Price(), time.Now(), time.Now())

	if err != nil {
		return err
	}

	defer instance.closeConnection(conn)

	return nil
}

func (instance RoomPostgresRepository) ListAllRooms() ([]room.Room, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	rows, err := conn.Query("SELECT * FROM room;")

	if err != nil {
		return nil, fmt.Errorf("falha ao obter usuários: %v", err)
	}

	var rooms []room.Room

	for rows.Next() {
		var id uuid.UUID
		var cod string
		var number int
		var vacancies int
		var price float64
		var createdAt time.Time
		var updatedAt time.Time

		err = rows.Scan(&id, &cod, &number, &vacancies, &price, &createdAt, &updatedAt)

		if err != nil {
			return nil, fmt.Errorf("falha ao obter usuários: %v", err)
		}

		newRoom, err := room.NewBuilder().WithID(id).WithCod(cod).WithNumber(number).WithVacancies(vacancies).WithPrice(price).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).Build()

		if err != nil {
			return nil, fmt.Errorf("falha ao obter usuários: %v", err)
		}

		rooms = append(rooms, *newRoom)

	}

	defer instance.closeConnection(conn)

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

func NewRoomPostgresRepository(connectorManager connectorManager) *RoomPostgresRepository {
	return &RoomPostgresRepository{
		connectorManager: connectorManager,
	}
}
