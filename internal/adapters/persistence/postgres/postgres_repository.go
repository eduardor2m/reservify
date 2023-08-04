package postgres

import (
	"github.com/jmoiron/sqlx"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectorManager)(nil)

type DatabaseConnectorManager struct{}

func (dcm *DatabaseConnectorManager) getConnection() (*sqlx.DB, error) {
	uri := "postgres://root:root@localhost:5432/reservify?sslmode=disable"

	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id uuid PRIMARY KEY,
			name varchar(255),
			email varchar(255) UNIQUE,
			password varchar(255),
			date_of_birth varchar(255),
			admin boolean
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (dcm *DatabaseConnectorManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()

	if err != nil {
		panic(err)
	}
}
