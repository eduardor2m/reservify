package postgres

import (
	"context"
	"fmt"
	"os"
	"reservify/internal/adapters/persistence/postgres/bridge"
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var _ repository.UserLoader = &UserPostgresRepository{}

type UserPostgresRepository struct {
	connectorManager
}

func checkIfUserIsAdmin(tokenJwt string, queries bridge.Queries, ctx context.Context) error {
	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenJwt[7:], func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return fmt.Errorf("falha ao obter token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(string)

		userIDFromToken, err := uuid.Parse(userID)

		if err != nil {
			return fmt.Errorf("falha ao converter id do usuário: %v", err)
		}

		userDB, err := queries.FindUserByID(ctx, userIDFromToken)

		if err != nil {
			return fmt.Errorf("falha ao encontrar usuário: %v", err)
		}

		if !userDB.Admin {
			return fmt.Errorf("usuário não é administrador")
		}

	} else {
		fmt.Println("Token inválido.")
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
		Admin:       false,
		CreatedAt:   u.CreatedAt(),
		UpdatedAt:   u.UpdatedAt(),
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
	claims["user_id"] = userDB.ID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return nil, fmt.Errorf("falha ao criar token: %v", err)
	}

	return &tokenString, nil
}

func (instance UserPostgresRepository) ListAllUsers(tokenJwt string) ([]user.User, error) {
	var users []user.User

	conn, err := instance.getConnection()

	if err != nil {
		return users, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	err = checkIfUserIsAdmin(tokenJwt, *queries, ctx)

	if err != nil {
		return users, err
	}

	usersDB, err := queries.ListAllUsers(ctx)

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

	userDB, err := queries.FindUserByID(ctx, id)

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

func (instance UserPostgresRepository) GetUsersByName(name string) ([]user.User, error) {
	var users []user.User

	conn, err := instance.getConnection()

	if err != nil {
		return users, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	queries := bridge.New(conn)

	ctx := context.Background()

	usersDB, err := queries.ListUsersByName(ctx, name)

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

	err = queries.UpdateUserByEmail(ctx, bridge.UpdateUserByEmailParams{
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

	err = queries.DeleteUserByEmail(ctx, email)

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
