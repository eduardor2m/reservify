package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectorManager)(nil)

type DatabaseConnectorManager struct{}

func (dcm *DatabaseConnectorManager) getConnection() (*sqlx.DB, error) {
	uri := "postgres://root:root@localhost:5432/reservify?sslmode=disable"

	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS "user" (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			date_of_birth VARCHAR(255),
			admin BOOLEAN,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
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
