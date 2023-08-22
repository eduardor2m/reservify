package postgres

import (
	"context"
	"fmt"
	"reservify/internal/adapters/persistence/postgres/bridge"
	"reservify/internal/app/entity/image"
	"reservify/internal/app/entity/room"
	"reservify/internal/app/interfaces/repository"
	"reservify/internal/utils/converters"

	"github.com/google/uuid"
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

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = queries.CreateRoom(ctx, bridge.CreateRoomParams{
		ID:        u.ID(),
		Cod:       u.Cod(),
		Number:    int32(u.Number()),
		Vacancies: int32(u.Vacancies()),
		Price:     converters.FloatToString(u.Price()),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	})

	if err != nil {
		return fmt.Errorf("falha ao criar quarto: %v", err)
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
		return nil, fmt.Errorf("falha ao obter quarto: %v", err)
	}

	var rooms []room.Room

	for _, roomDB := range roomsDB {
		price, err := converters.StringToFloat(roomDB.Price)

		if err != nil {
			return nil, fmt.Errorf("falha ao converter preço do quarto: %v", err)
		}

		roomBuild, err := room.NewBuilder().WithID(roomDB.ID).WithCod(roomDB.Cod).WithNumber(int(roomDB.Number)).WithVacancies(int(roomDB.Vacancies)).WithPrice(price).WithCreatedAt(roomDB.CreatedAt).WithUpdatedAt(roomDB.UpdatedAt).Build()

		if err != nil {
			return nil, fmt.Errorf("falha ao construir quarto: %v", err)
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

	queries := bridge.New(conn)

	ctx := context.Background()

	roomDB, err := queries.FindRoomById(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter quarto: %v", err)
	}

	price, err := converters.StringToFloat(roomDB.Price)

	if err != nil {
		return nil, fmt.Errorf("falha ao converter preço do quarto: %v", err)
	}

	roomBuild, err := room.NewBuilder().WithID(roomDB.ID).WithCod(roomDB.Cod).WithNumber(int(roomDB.Number)).WithVacancies(int(roomDB.Vacancies)).WithPrice(price).WithCreatedAt(roomDB.CreatedAt).WithUpdatedAt(roomDB.UpdatedAt).Build()

	if err != nil {
		return nil, fmt.Errorf("falha ao construir quarto: %v", err)
	}

	return roomBuild, nil
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

	price, err := converters.StringToFloat(roomDB.Price)

	if err != nil {
		return nil, fmt.Errorf("falha ao converter preço do quarto: %v", err)
	}

	roomBuild, err := room.NewBuilder().WithID(roomDB.ID).WithCod(roomDB.Cod).WithNumber(int(roomDB.Number)).WithVacancies(int(roomDB.Vacancies)).WithPrice(price).WithCreatedAt(roomDB.CreatedAt).WithUpdatedAt(roomDB.UpdatedAt).Build()

	if err != nil {
		return nil, fmt.Errorf("falha ao construir quarto: %v", err)
	}

	return roomBuild, nil

}

func (instance RoomPostgresRepository) DeleteRoomByID(id uuid.UUID) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = queries.DeleteRoomById(ctx, id)

	if err != nil {
		return fmt.Errorf("falha ao deletar quarto: %v", err)
	}

	return nil
}

func (instance RoomPostgresRepository) AddImageToRoomById(id uuid.UUID, image string) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = queries.AddImageToRoomByID(ctx,
		bridge.AddImageToRoomByIDParams{
			IDRoom:   id,
			ImageUrl: image,
		},
	)

	if err != nil {
		return fmt.Errorf("falha ao adicionar imagem ao quarto: %v", err)
	}

	return nil
}

func (instance RoomPostgresRepository) ListAllImagesByRoomID(id uuid.UUID) (*room.Room, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	imagesDB, err := queries.ListAllImagesByRoomID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter imagens do quarto: %v", err)
	}

	roomDB, err := queries.FindRoomById(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter quarto: %v", err)
	}

	price, err := converters.StringToFloat(roomDB.Price)

	if err != nil {
		return nil, fmt.Errorf("falha ao converter preço do quarto: %v", err)
	}

	roomBuild, err := room.NewBuilder().WithID(roomDB.ID).WithCod(roomDB.Cod).WithNumber(int(roomDB.Number)).WithVacancies(int(roomDB.Vacancies)).WithPrice(price).WithCreatedAt(roomDB.CreatedAt).WithUpdatedAt(roomDB.UpdatedAt).Build()

	if err != nil {
		return nil, fmt.Errorf("falha ao construir quarto: %v", err)
	}

	var imagesBuild []image.Image

	for _, imageDb := range imagesDB {
		imageBuild, err := image.NewBuilder().WithIDRoom(imageDb.IDRoom).WithImageUrl(imageDb.ImageUrl).Build()

		if err != nil {
			return nil, fmt.Errorf("falha ao construir imagem: %v", err)
		}

		imagesBuild = append(imagesBuild, *imageBuild)
	}

	roomBuild.Image = imagesBuild

	return roomBuild, nil
}

func NewRoomPostgresRepository(connectorManager connectorManager) *RoomPostgresRepository {
	return &RoomPostgresRepository{
		connectorManager: connectorManager,
	}
}
