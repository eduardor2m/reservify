package postgres

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"reservify/internal/adapters/persistence/postgres/bridge"
	"reservify/internal/app/entity/user"
	"reservify/internal/app/interfaces/repository"
	"time"
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
		DateOfBirth: time.Now(),
		Admin:       u.Admin(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %v", err)
	}

	return nil
}

func (instance UserPostgresRepository) LoginUser(email string, password string) (error, *string) {
	conn, err := instance.getConnection()
	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err), nil
	}
	defer instance.closeConnection(conn)

	var userPass string

	err = conn.QueryRow(`
		SELECT password FROM "user" WHERE email = $1;
	`, email).Scan(&userPass)

	err = bcrypt.CompareHashAndPassword([]byte(userPass), []byte(password))

	if err != nil {
		return fmt.Errorf("falha ao logar usuário: %v", err), nil
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return fmt.Errorf("falha ao logar usuário: %v", err), nil
	}

	return nil, &tokenString
}

func (instance UserPostgresRepository) RentRoom(
	idUser string,
	idRoom string,
	checkIn string,
	checkOut string,
) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

	defer instance.closeConnection(conn)

	newReservationUUID, err := uuid.NewUUID()

	_, err = conn.Exec(`
		INSERT INTO reservation (id, id_user, id_room, check_in, check_out) VALUES ($1, $2, $3, $4, $5);
	`, newReservationUUID, idUser, idRoom, checkIn, checkOut)

	if err != nil {
		return fmt.Errorf("falha ao alugar quarto: %v", err)
	}

	return nil
}

func (instance UserPostgresRepository) ListAllUsers() ([]user.User, error) {
	var users []user.User

	conn, err := instance.getConnection()

	defer instance.closeConnection(conn)

	if err != nil {
		return users, fmt.Errorf("falha ao obter conexão com o banco de dados: %v", err)
	}

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
			WithDateOfBirth(userDB.DateOfBirth.String()).
			WithAdmin(userDB.Admin).
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

	var name, email, password, dateOfBirth, cpf, phone string
	var createdAt, updatedAt time.Time
	var admin bool
	var idUser uuid.UUID

	err = conn.QueryRow(`
		SELECT id, name, email, password, cpf, phone, date_of_birth, admin, created_at, updated_at FROM "user" WHERE id = $1;
	`, id).Scan(&idUser, &name, &email, &password, &cpf, &phone, &dateOfBirth, &admin, &createdAt, &updatedAt)

	if err != nil {
		return nil, fmt.Errorf("falha ao obter usuário: %v", err)
	}

	userReceived, err := user.NewBuilder().
		WithID(idUser).
		WithName(name).
		WithEmail(email).
		WithCPF(cpf).
		WithPhone(phone).
		WithPassword(password).
		WithDateOfBirth(dateOfBirth).
		WithAdmin(admin).
		WithCreatedAt(createdAt).
		WithUpdatedAt(updatedAt).
		Build()

	if err != nil {
		return nil, fmt.Errorf("falha ao construir usuário: %v", err)
	}

	return userReceived, nil
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
	fmt.Println(user, email)

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
