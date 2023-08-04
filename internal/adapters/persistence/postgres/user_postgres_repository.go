package postgres

import (
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/repository"
)

var _ repository.UserLoader = &UserPostgresRepository{}

type UserPostgresRepository struct {
	connectorManager
}

func (instance UserPostgresRepository) CreateUser(user user.User) error {
	conn, err := instance.getConnection()

	if err != nil {
		return err
	}

	defer instance.closeConnection(conn)

	_, err = conn.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (instance UserPostgresRepository) GetUserByName(name string) ([]user.User, error) {
	var users []user.User

	return users, nil
}

func (instance UserPostgresRepository) UpdateUserByEmail(email string, user user.User) error {

	return nil
}

func (instance UserPostgresRepository) DeleteUserByEmail(email string) error {

	return nil
}

func NewUserPostgresRepository(connectorManager connectorManager) *UserPostgresRepository {
	return &UserPostgresRepository{
		connectorManager: connectorManager,
	}
}
