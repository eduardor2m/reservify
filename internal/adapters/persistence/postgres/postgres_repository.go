package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectorManager)(nil)

type DatabaseConnectorManager struct{}

func RunMigrations() error {
	var (
		dbUser     string
		dbPassword string
		dbHost     string
		dbPort     string
		dbName     string
	)

	development := os.Getenv("DEVELOPMENT")

	if development == "true" {
		dbUser = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	} else {
		dbUser = os.Getenv("DB_USER_GCP")
		dbPassword = os.Getenv("DB_PASSWORD_GCP")
		dbHost = os.Getenv("DB_HOST_GCP")
		dbPort = os.Getenv("DB_PORT_GCP")
		dbName = os.Getenv("DB_NAME_GCP")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS "user" (
		id UUID PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		cpf VARCHAR(11) NOT NULL UNIQUE,
		email VARCHAR(255) NOT NULL UNIQUE,
		phone VARCHAR(50) NOT NULL,
		date_of_birth DATE NOT NULL,
		password VARCHAR(255) NOT NULL,
		admin BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS room (
		id UUID PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		cod VARCHAR(255) NOT NULL UNIQUE,
		number INTEGER NOT NULL,
		vacancies INTEGER NOT NULL,
		price DECIMAL(10,2) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS "image" (
		id_room UUID NOT NULL,
		image_url TEXT PRIMARY KEY,
		thumbnail BOOLEAN NOT NULL DEFAULT FALSE,
		FOREIGN KEY (id_room) REFERENCES room(id) ON DELETE CASCADE
	);
`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS "reservation" (
		id UUID PRIMARY KEY,
		id_user UUID NOT NULL,
		id_room UUID NOT NULL,
		check_in VARCHAR(10) NOT NULL,
		check_out VARCHAR(10) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		FOREIGN KEY (id_user) REFERENCES "user"(id),
		FOREIGN KEY (id_room) REFERENCES room(id)
	);
`)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return nil
}

func (dcm *DatabaseConnectorManager) getConnection() (*sqlx.DB, error) {
	var (
		dbUser     string
		dbPassword string
		dbHost     string
		dbPort     string
		dbName     string
	)

	development := os.Getenv("DEVELOPMENT")

	if development == "true" {
		dbUser = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	} else {
		dbUser = os.Getenv("DB_USER_GCP")
		dbPassword = os.Getenv("DB_PASSWORD_GCP")
		dbHost = os.Getenv("DB_HOST_GCP")
		dbPort = os.Getenv("DB_PORT_GCP")
		dbName = os.Getenv("DB_NAME_GCP")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return db, nil
}

func (dcm *DatabaseConnectorManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()

	if err != nil {
		panic(err)
	}
}
