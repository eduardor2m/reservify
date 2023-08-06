package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/repository"
)

var _ repository.UserLoader = &UserPostgresRepository{}

type UserPostgresRepository struct {
	connectorManager
}

func (instance UserPostgresRepository) CreateUser(u user.User) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	_, err = conn.Exec(`
		INSERT INTO "user" (id, name, email, password, date_of_birth, admin) VALUES ($1, $2, $3, $4, $5, $6);
	`, u.ID(), u.Name(), u.Email(), u.Password(), u.DateOfBirth(), u.Admin())

	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %v", err)
	}

	defer instance.closeConnection(conn)

	return nil
}

func (instance UserPostgresRepository) ListAllUsers() ([]user.User, error) {
	var users []user.User

	conn, err := instance.getConnection()

	defer instance.closeConnection(conn)

	if err != nil {
		return users, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	query := `
		SELECT id, name, email, password, date_of_birth, admin, created_at, updated_at FROM "user";
	`

	rows, err := conn.Query(query)
	if err != nil {
		return users, fmt.Errorf("falha ao listar usuários: %v", err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	for rows.Next() {
		var id uuid.UUID
		var name, email, password, dateOfBirth, createdAt, updatedAt string
		var admin bool

		err := rows.Scan(&id, &name, &email, &password, &dateOfBirth, &admin, &createdAt, &updatedAt)
		if err != nil {
			return users, fmt.Errorf("falha ao listar usuários: %v", err)
		}

		userReceived, err := user.NewBuilder().
			WithID(id).
			WithName(name).
			WithEmail(email).
			WithPassword(password).
			WithDateOfBirth(dateOfBirth).
			WithAdmin(admin).
			Build()

		if err != nil {
			return users, fmt.Errorf("falha ao construir usuário: %v", err)
		}

		users = append(users, *userReceived)
	}

	return users, nil
}

func (instance UserPostgresRepository) GetUserByName(name string) ([]user.User, error) {
	var users []user.User

	conn, err := instance.getConnection()

	defer instance.closeConnection(conn)

	if err != nil {
		return users, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	query := `
		SELECT id, name, email, password, date_of_birth, admin, created_at, updated_at FROM "user" WHERE name = $1;
	`

	rows, err := conn.Query(query, name)

	if err != nil {
		return users, fmt.Errorf("falha ao listar usuários: %v", err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	for rows.Next() {
		var id uuid.UUID
		var name, email, password, dateOfBirth, createdAt, updatedAt string
		var admin bool

		err := rows.Scan(&id, &name, &email, &password, &dateOfBirth, &admin, &createdAt, &updatedAt)
		if err != nil {
			return users, fmt.Errorf("falha ao listar usuários: %v", err)
		}

		userReceived, err := user.NewBuilder().
			WithID(id).
			WithName(name).
			WithEmail(email).
			WithPassword(password).
			WithDateOfBirth(dateOfBirth).
			WithAdmin(admin).
			Build()

		if err != nil {
			return users, fmt.Errorf("falha ao construir usuário: %v", err)
		}

		users = append(users, *userReceived)
	}

	return users, nil
}

func (instance UserPostgresRepository) UpdateUserByEmail(email string, user user.User) error {

	return nil
}

func (instance UserPostgresRepository) DeleteUserByEmail(email string) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	_, err = conn.Exec(`
		DELETE FROM "user" WHERE email = $1;
	`, email)

	if err != nil {
		return fmt.Errorf("falha ao deletar usuário: %v", err)
	}

	defer instance.closeConnection(conn)

	return nil

}

func NewUserPostgresRepository(connectorManager connectorManager) *UserPostgresRepository {
	return &UserPostgresRepository{
		connectorManager: connectorManager,
	}
}
