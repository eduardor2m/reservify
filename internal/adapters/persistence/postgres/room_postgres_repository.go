package postgres

import (
	"fmt"
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

	stmt, err := conn.Prepare("INSERT INTO room (id, cod, number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)")

	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %v", err)
	}

	_, err = stmt.Exec(u.ID(), u.Cod(), u.Number(), time.Now(), time.Now())

	if err != nil {
		return err
	}

	defer instance.closeConnection(conn)

	return nil
}

func NewRoomPostgresRepository(connectorManager connectorManager) *RoomPostgresRepository {
	return &RoomPostgresRepository{
		connectorManager: connectorManager,
	}
}
